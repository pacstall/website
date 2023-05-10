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
                    ' Pacstall takes in files known as <0>pacscripts</0> (similar' +
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
    packageSearch: {
        dropdown: {
            package: 'Package',
            packageTooltip: 'Searches in package names and descriptions',
            maintainer: 'Maintainer',
            maintainerTooltip: 'Searches by maintainer names and emails',
        },
        table: {
            name: 'Name',
            maintainer: 'Maintainer',
            version: 'Version',
            install: 'Install',
        },
        versionTooltip: {
            notInRegistry: 'This package is not in the Repology registry',
            latest: 'This package is the latest version',
            hasPatchUpdate: 'This package has a patch update available',
            hasMinorUpdate: 'This package has a minor update available',
            hasMajorUpdate: 'This package has a major update available',
            isGit: 'This package is a Git package',
        },
        noResults: 'Not finding what you want? <0>Create a request!</0>',
        search: 'Search',
        orphaned: 'Orphaned',
        maintainerTooltip: {
            maintainedBy: 'This package is being maintained by {{ name }}',
            noMaintainer: 'This package is not being maintained',
        },
        pagination: {
            previous: 'previous',
            next: 'next',
        },
    },
    packageDetails: {
        table: {
            name: 'Name',
            version: 'Version',
            maintainer: 'Maintainer',
            dependencies: 'Dependencies',
            requiredBy: 'Required By',
        },
        orphaned: 'Orphaned',
        noResults: 'None',
        openInGithub: 'Open In Github',
        view: 'View',
        howToInstall: {
            title: 'How to Install',
            step1: 'Step 1: Setup Pacstall',
            step2: 'Step 2: Install {{ name }}',
        },
        dependenciesModal: {
            title: 'Dependencies',
            buildDependencies: 'Build Dependencies',
            optionalDependencies: 'Optional Dependencies',
            runtimeDependencies: 'Runtime Dependencies',
            pacstallDependencies: 'Pacstall Dependencies',
            name: 'Name',
            close: 'Close',
            provider: 'Provider',
            noDescription: 'No description available',
        },
        requiredByModal: {
            title: 'Required By',
            name: 'Name',
            provider: 'Provider',
            close: 'Close',
            noDescription: 'No description available',
        },
    },
}
