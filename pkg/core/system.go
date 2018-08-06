/*
 * Copyright 2018 The original author or authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package core

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
	"bufio"
	"os"
	"strings"
)

const (
	istioNamespace  = "istio-system"
	istioRelease    = "https://storage.googleapis.com/riff-releases/istio-riff-0.1.1.yaml"
	servingRelease  = "https://storage.googleapis.com/riff-releases/serving-release-no-mon-riff-0.1.1.yaml"
	eventingRelease = "https://storage.googleapis.com/riff-releases/eventing-release-riff-0.1.1.yaml"
	stubBusRelease  = "https://storage.googleapis.com/riff-releases/eventing-release-clusterbus-stub-riff-0.1.1.yaml"
)

type SystemInstallOptions struct {
	NodePort bool
	Force bool
}

type SystemUninstallOptions struct {
	Istio bool
	Force bool
}

var (
	knativeNameSpaces = []string{"knative-eventing", "knative-serving", "knative-build"}
	allNameSpaces = append(knativeNameSpaces, istioNamespace)
)

func (kc *kubectlClient) SystemInstall(options SystemInstallOptions) error {

	err := ensureNotTerminating(kc, allNameSpaces, "Please try again later.")
	if err != nil {
		return err
	}

	istioStatus, err := getNamespaceStatus(kc,istioNamespace)
	if istioStatus == "'NotFound'" {
		fmt.Print("Installing Istio Components\n")
		istioYaml, err := loadRelease(istioRelease)
		if err != nil {
			return err
		}
		if options.NodePort {
			istioYaml = bytes.Replace(istioYaml, []byte("LoadBalancer"), []byte("NodePort"), -1)
		}
		fmt.Printf("Applying resources defined in: %s\n", istioRelease)
		istioLog, err := kc.kubeCtl.ExecStdin([]string{"apply", "-f", "-"}, &istioYaml)
		if err != nil {
			fmt.Printf("%s\n", istioLog)
			return err
		}

		fmt.Print("Istio for riff installed\n\n")
	} else {
		if !options.Force {
			answer, err := confirm("Istio is already installed, do you want to install the Knative components for riff?")
			if err != nil {
				return err
			}
			if !answer {
				return nil
			}
		}
	}

	err = waitForIstioSidecarInjector(kc)
	if err != nil {
		return err
	}

	fmt.Print("Installing Knative Components\n")

	servingYaml, err := loadRelease(servingRelease)
	if err != nil {
		return err
	}
	if options.NodePort {
		servingYaml = bytes.Replace(servingYaml, []byte("LoadBalancer"), []byte("NodePort"), -1)
	}
	fmt.Printf("Applying resources defined in: %s\n", servingRelease)
	servingLog, err := kc.kubeCtl.ExecStdin([]string{"apply", "-f", "-"}, &servingYaml)
	if err != nil {
		fmt.Printf("%s\n", servingLog)
		return err
	}

	applyResources(kc, eventingRelease)

	applyResources(kc, stubBusRelease)

	fmt.Print("Knative for riff installed\n\n")
	return nil
}

func (kc *kubectlClient) SystemUninstall(options SystemUninstallOptions) (error, bool) {

	err := ensureNotTerminating(kc, allNameSpaces, "This would indicate that the system was already uninstalled.")
	if err != nil {
		return err, false
	}
	knativeNsCount, err := checkNamespacesExists(kc, knativeNameSpaces)
	istioNsCount, err := checkNamespacesExists(kc, []string{istioNamespace})
	if err != nil {
		return err, false
	}
	if knativeNsCount == 0 {
		fmt.Print("No Knative components for riff found\n")
	} else {
		if !options.Force {
			answer, err := confirm("Are you sure you want to uninstall the riff system?")
			if err != nil {
				return err, false
			}
			if !answer {
				return nil, false
			}
		}
		fmt.Print("Removing Knative for riff components\n")
		err = deleteCrds(kc, "knative.dev")
		if err != nil {
			return err, false
		}
		err = deleteClusterResources(kc, "clusterrolebinding", "knative-")
		if err != nil {
			return err, false
		}
		err = deleteClusterResources(kc, "clusterrolebinding", "build-controller-")
		if err != nil {
			return err, false
		}
		err = deleteClusterResources(kc, "clusterrolebinding", "eventing-controller-")
		if err != nil {
			return err, false
		}
		err = deleteClusterResources(kc, "clusterrolebinding", "clusterbus-controller-")
		if err != nil {
			return err, false
		}
		err = deleteClusterResources(kc, "clusterrole", "knative-")
		if err != nil {
			return err, false
		}
		err = deleteNamespaces(kc, knativeNameSpaces)
		if err != nil {
			return err, false
		}
	}
	if istioNsCount == 0 {
		fmt.Print("No Istio components found\n")
	} else {
		if !options.Istio {
			if options.Force {
				return nil, true
			}
			answer, err := confirm("Do you also want to uninstall Istio components?")
			if err != nil {
				return err, false
			}
			if !answer {
				return nil, true
			}
		}
		fmt.Print("Removing Istio components\n")
		err = deleteCrds(kc, "istio.io")
		if err != nil {
			return err, false
		}
		err = deleteClusterResources(kc, "clusterrolebinding", "istio-")
		if err != nil {
			return err, false
		}
		err = deleteClusterResources(kc, "clusterrole", "istio-")
		if err != nil {
			return err, false
		}
		err = deleteNamespaces(kc, []string{istioNamespace})
		if err != nil {
			return err, false
		}
	}
	return nil, true
}

func resolveReleaseURLs(filename string) (url.URL, error) {
	u, err := url.Parse(filename)
	if err != nil {
		return url.URL{}, err
	}
	if u.Scheme == "http" || u.Scheme == "https" {
		return *u, nil
	}
	return *u, fmt.Errorf("filename must be file, http or https, got %s", u.Scheme)
}

func loadRelease(release string) ([]byte, error) {
	releaseUrl, err := resolveReleaseURLs(release)
	if err != nil {
		return nil, err
	}
	resp, err := http.Get(releaseUrl.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func waitForIstioSidecarInjector(kc *kubectlClient) error {
	print("Waiting for istio-sidecar-injector to start ")
	for i := 0; i < 36; i++ {
		fmt.Print(".")
		injectorStatus, err := kc.kubeCtl.Exec([]string{"get", "pod", "-n", istioNamespace, "-l", "istio=sidecar-injector", "-o", "jsonpath='{.items[0].status.phase}'"})
		if err != nil {
			// might take some time for the pod to show up so ignore early errors
			if i > 3 {
				return err
			}
		}
		if injectorStatus == "'Error'" {
			return errors.New("istio-sidecar-injector pod failed to start")
		}
		if injectorStatus == "'Running'" {
			print(injectorStatus, "\n\n")
			return nil
		}
		time.Sleep(10 * time.Second) // wait for it to start
	}
	fmt.Print("\n\n")
	return errors.New("istio-sidecar-injector pod did not start in time")
}

func applyResources(kc *kubectlClient, release string) error {
	releaseUrl, err := resolveReleaseURLs(release)
	if err != nil {
		return err
	}
	fmt.Printf("Applying resources defined in: %s\n", releaseUrl.String())
	releaseLog, err := kc.kubeCtl.Exec([]string{"apply", "-f", releaseUrl.String()})
	if err != nil {
		fmt.Printf("%s", releaseLog)
	}
	return nil
}

func deleteNamespaces(kc *kubectlClient, namespaces []string) error {
	for _, namespace := range namespaces {
		fmt.Printf("Deleting resources defined in: %s\n", namespace)
		deleteLog, err := kc.kubeCtl.Exec([]string{"delete", "namespace", namespace})
		if err != nil {
			fmt.Printf("%s", deleteLog)
		}
	}
	return nil
}

func deleteClusterResources(kc *kubectlClient, resourceType string, prefix string) error {
	fmt.Printf("Deleting %ss prefixed with %s\n", resourceType, prefix)
	resourceList, err := kc.kubeCtl.Exec([]string{"get", resourceType, "-ocustom-columns=name:metadata.name"})
	if err != nil {
		return err
	}
	resource := strings.Split(string(resourceList), "\n")
	var resourcesToDelete []string
	for _, resource := range resource {
		if strings.HasPrefix(resource, prefix) {
			resourcesToDelete = append(resourcesToDelete, resource)
		}
	}
	if len(resourcesToDelete) > 0 {
		resourceLog, err := kc.kubeCtl.Exec(append([]string{"delete", resourceType}, resourcesToDelete...))
		if err != nil {
			fmt.Printf("%s", resourceLog)
			return err
		}
	}
	return nil
}

func deleteCrds(kc *kubectlClient, suffix string) error {
	fmt.Printf("Deleting CRDs for %s\n", suffix)
	crdList, err := kc.kubeCtl.Exec([]string{"get", "customresourcedefinitions", "-ocustom-columns=name:metadata.name"})
	if err != nil {
		return err
	}
	crds := strings.Split(string(crdList), "\n")
	var crdsToDelete []string
	for _, crd := range crds {
		if strings.HasSuffix(crd, suffix) {
			crdsToDelete = append(crdsToDelete, crd)
		}
	}
	if len(crdsToDelete) > 0 {
		crdLog, err := kc.kubeCtl.Exec(append([]string{"delete", "customresourcedefinition"}, crdsToDelete...))
		if err != nil {
			fmt.Printf("%s", crdLog)
			return err
		}
	}
	return nil
}

func checkNamespacesExists (kc *kubectlClient, names []string) (int, error) {
	count := 0
	for _, name := range names {
		status, err := getNamespaceStatus(kc, name)
		if err != nil {
			return count, err
		}
		if status != "'NotFound'" {
			count =+ 1
		}
	}
	return count, nil
}

func ensureNotTerminating (kc *kubectlClient, names []string, message string) error {
	for _, name := range names {
		status, err := getNamespaceStatus(kc, name)
		if err != nil {
			return err
		}
		if status == "'Terminating'" {
			return errors.New(fmt.Sprintf("The %s namespace is currently 'Terminating'. %s", name, message))
		}
	}
	return nil
}

func getNamespaceStatus(kc *kubectlClient, name string) (string, error) {
	nsLog, err := kc.kubeCtl.Exec([]string{"get", "namespace", name, "-o", "jsonpath='{.status.phase}'"})
	if err != nil {
		if strings.Contains(nsLog, "NotFound") {
			return "'NotFound'", nil
		}
		return "", err
	}
	return nsLog, nil
}

func confirm(s string) (bool, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("%s [y/n]: ", s)
	res, err := reader.ReadString('\n')
	if err != nil {
		return false, err
	}
	if len(res) < 2 {
		return false, nil
	}
	answer := strings.ToLower(strings.TrimSpace(res))[0] == 'y'
	return answer, nil
}