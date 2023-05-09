import Locale from './locale'

export default <Locale>{
    home: {
        title: 'Pacstall',
        subtitle: 'The AUR for Ubuntu',
        cards: {
            whyDifferent: {
                title: 'Why is this any different than any other package manager?',
                description:
                    'Pacstall uses the stable base of Ubuntu but allows you to use ' +
                    "bleeding edge software with little to no compromises, so you don't " +
                    'have to worry about security patches or new features.',
            },
            howItWorks: {
                title: 'How does it work then?',
                description:
                    ' Pacstall takes in files known as <link>pacscripts</link> (similar' +
                    ' to PKGBUILDs) that contain the necessary contents to build packages,' +
                    ' and builds them into executables on your system.',
            },
        },
        installationInstructions: 'Installation Instructions',
        showcase: {
            title: 'Showcase',
            packageSearch: 'Package Search',
        },
    },
    navbar: {
        title: 'Pacstall',
        contribute: {
            title: 'Contribute',
            workOnFeatures: 'Work on new features',
            helpTranslate: 'Help with translations',
            becomeAMaintainer: 'Become a package maintainer',
        },
        social: {
            title: 'Social Networks',
            discord: 'Discord',
            matrix: 'Matrix',
            reddit: 'Reddit',
            mastodon: 'Mastodon',
        },
        browse: {
            title: 'Browse Packages',
        },
        privacy: {
            title: 'Privacy Policy',
        },
        install: 'Install',
    },
    cookieConsent: {
        title: 'Cookie notice TL;DR',
        paragraphs: [
            'Hi, yes, we use cookies.',
            "We don't like to give you misleading or confusing information. We only use cookies for essential features such as theme settings, localization and authentication.",
            'You can read the full privacy policy <0>here <1/></0>.',
            'By continuing to use this site, you <strong>give your agreement to the privacy policy</strong>.',
        ],
        accept: 'ok, nice',
    },
}
