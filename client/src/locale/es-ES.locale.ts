import Locale from './locale'

export default <Locale>{
    home: {
        title: 'Pacstall',
        subtitle: 'El AUR para Ubuntu',
        cards: {
            whyDifferent: {
                title: '¿Por qué es diferente a cualquier otro administrador de paquetes?',
                description:
                    'Pacstall usa la base estable de Ubuntu pero le permite ' +
                    'usar software de última generación con poco o ningún compromiso, ' +
                    'por lo que no tiene que preocuparse por los parches de seguridad ' +
                    'o las nuevas funciones.',
            },
            howItWorks: {
                title: '¿Cómo funciona entonces?',
                description:
                    ' Pacstall takes in files known as <link>pacscripts</link> (similar' +
                    ' to PKGBUILDs) that contain the necessary contents to build packages,' +
                    ' and builds them into executables on your system.',
            },
        },
        installationInstructions: 'Instrucciones de Instalación',
        showcase: {
            title: 'Escaparate',
            packageSearch: 'Búsqueda de Paquetes',
        },
    },
    navbar: {
        title: 'Pacstall',
        contribute: {
            title: 'Contribuir',
            workOnFeatures: 'Trabaja en nuevas características',
            helpTranslate: 'Ayuda con las traducciones',
            becomeAMaintainer: 'Conviértase en un mantenedor de paquetes',
        },
        social: {
            title: 'Redes Sociales',
            discord: 'Discord',
            matrix: 'Matrix',
            reddit: 'Reddit',
            mastodon: 'Mastodon',
        },
        browse: {
            title: 'Explorar Paquetes',
        },
        privacy: {
            title: 'La Política de Privacidad',
        },
        install: 'Instalar',
    },
    cookieConsent: {
        title: 'Cookie notice TL;DR',
        paragraphs: [
            'Hi, yes, we use cookies.',
            "We don't like to give you misleading or confusing information. We only use cookies for essential features such as theme settings, localization and authentication.",
            'You can read the full privacy policy <0>here <1/></0>',
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
            install: 'Instalar',
        },
        versionTooltip: {
            notInRegistry: 'This package is not in the Repology registry',
            latest: 'This package is the latest version',
            hasPatchUpdate: 'This package has a patch update available',
            hasMinorUpdate: 'This package has a minor update available',
            hasMajorUpdate: 'This package has a major update available',
            isGit: 'This package is a Git package',
        },
        orphaned: 'Orphaned',
        noResults: 'Not finding what you want? <0>Create a request!</0>',
        search: 'Search',
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
