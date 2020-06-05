const siteConfig = {
  title: 'Wspieram przyrodę ',
  tagline: 'Dokumentacja',
  url: 'https://docs.wspieramprzyrode.pl',
  baseUrl: '/',
  projectName: 'app',
  organizationName: 'wspieramprzyrode',
  headerLinks: [
    { href: "https://github.com/wspieramprzyrode/app", label: "GitHub App" },
    { href: "https://github.com/wspieramprzyrode/mobile", label: "GitHub Mobile" },
  ],
  disableHeaderTitle: true,
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
