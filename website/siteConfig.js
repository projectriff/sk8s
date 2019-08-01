// https://docusaurus.io/docs/site-config

// This site does not use built-in docusuarus versioning or docusaurus publish
// projectriff.io is generated by running a docusarus build per version
// See /website/README.md for details about adding versions

const siteConfig = {

  projectName: 'riff',
  organizationName: 'projectriff',
  title: 'riff is for functions',
  tagline: '',

  cname: 'projectriff.io',
  url: 'https://projectriff.io',
  baseUrl: '/',

  // version generated by this build
  docsUrl: 'docs/0-4',

  // used only in pages/en/versions.js
  versions: [
    { name:'v0.4.x (next)', url:'docs/0-4' },
    { name:'v0.3.x (stable)', url:'docs/0-3' }
  ],

  headerLinks: [
    {page: 'versions', label: 'v0.4.x'},
    {doc: 'getting-started', label: 'Docs'},
    {doc: 'riff', label: 'CLI'},
    {blog: true, label: 'Blog'},
  ],

  headerIcon: 'img/riff-white.svg',
  footerIcon: 'img/riff-white.svg',
  favicon: 'img/favicon.ico',

  colors: {
    primaryColor: '#52adc8',
    secondaryColor: '#111111',
  },

  // theme for syntax highlighting
  highlight: {
    theme: 'default',
  },

  // on-page navigation
  onPageNav: 'separate',

  // no .html extensions
  cleanUrl: true,

  // open Graph and twitter card images
  ogImage: 'img/riff.svg',
  twitterImage: 'img/riff.svg',
  
  // show all blog posts in sidebar
  blogSidebarCount: 'ALL',

};

module.exports = siteConfig;
