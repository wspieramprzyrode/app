const siteConfig = {
  title: 'docs.wspieramprzyrode.pl ',
  tagline: 'Application Documentation',
  url: 'https://docs.wspieramprzyrode.pl',
  baseUrl: '/',
  projectName: 'app',
  organizationName: 'wspieramprzyrode',
  headerLinks: [
    { languages: true },
  ],
  translationRecruitingLink: 'https://crowdin.com/project/wspieramprzyrode',
  disableHeaderTitle: true,
  headerIcon: 'img/favicon.ico',
  footerIcon: 'img/favicon.ico',
  favicon: 'img/favicon.ico',
  colors: {
    primaryColor: "#2F1666",
    secondaryColor: "#7731F6"
  },
  copyright: `Copyright © ${new Date().getFullYear()} WspieramPrzyrode.pl`,

  highlight: {
    theme: 'atom-one-dark',
  },
  scripts: ['https://buttons.github.io/buttons.js'],
  onPageNav: 'separate',
  cleanUrl: true,
};

module.exports = siteConfig;
