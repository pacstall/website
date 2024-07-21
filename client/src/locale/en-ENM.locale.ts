import Locale from './locale'

export default <Locale>{
    home: {
        title: 'Pacstall',
        subtitle: 'Þe AUR for Ubuntu',
        cards: {
            whyDifferent: {
                title: 'Why is þis ani different þan any oþer pakagge manager?',
                description:
                    'Pacstall noteþ þe stable base of Ubuntu but alloweþ þe to notest ' +
                    'bledynge egge softeware wiþ litel to no compromises, so þou ne dost ' +
                    'have to troubleþ þeself wiþ securite pacches or newe fetures.',
            },
            howItWorks: {
                title: 'How þenne doþ hit werk?',
                description:
                    ' Pacstall takeþ in fyles witen as <0>pacscripts</0> (like' +
                    ' PKGBUILDs) þat holden þe necessarye contents to bylden pakagges,' +
                    ' and byldeþ þem in-to executables on þi systeme.',
            },
        },
        installationInstructions: 'Installacioun Instrucciouns',
        showcase: {
            title: 'Schewecas',
            packageSearch: 'Pakagge Serchen',
        },
    },
    navbar: {
        title: 'Pacstall',
        contribute: {
            title: 'Contribute',
            workOnFeatures: 'Werken on newe fetures',
            helpTranslate: 'Helpen wiþ translaciouns',
            becomeAMaintainer: 'Becomen a pakagge mayntener',
        },
        social: {
            title: 'Social Nettwerks',
            discord: 'Discord',
            matrix: 'Matrix',
            reddit: 'Reddit',
            lemmy: 'Lemmy',
            mastodon: 'Mastodon',
        },
        browse: {
            title: 'Browsen Pakagges',
        },
        privacy: {
            title: 'Privacie Policie',
        },
        install: 'Installen',
    },
    cookieConsent: {
        title: 'Cookie notice TL;DR',
        paragraphs: [
            'Hail, we are notynge cookies on ure website.',
            'We ne doþ wisshe to ȝiven ȝe misledynge ne mistakynge informacion. We oonly noten cookies for essential fetures swich as teme setynges, localisacioun and auþentikacioun.',
            'Ȝe can rede þe ful privacie policie <0>her <1/></0>.',
            'By continuenge to noten þis site, ȝe <strong>ȝivest eower agrement to þe privacie policie</strong>.',
        ],
        accept: "God be wi' ye",
    },
    packageSearch: {
        dropdown: {
            package: 'Pakagge',
            packageTooltip: 'Serchen in pakagge names and descripciouns',
            maintainer: 'Mayntener',
            maintainerTooltip: 'Serchen by mayntener names and emailes',
        },
        table: {
            name: 'Name',
            maintainer: 'Mayntener',
            version: 'Version',
            install: 'Install',
        },
        versionTooltip: {
            notInRegistry: 'Þis pakagge ne is in þe Repologgye reggistrye',
            latest: 'Þis pakagge is þe latest version',
            patch: 'Þis pakagge haveþ a patche update available',
            minor: 'Þis pakagge haveþ a minor update available',
            major: 'Þis pakagge haveþ a major update available',
            isGit: 'Þis pakagge is a Git pakagge',
        },
        noResults: 'Ne findynge what ȝe wisshest? <0>Create a request!</0>',
        search: 'Serchen',
        orphaned: 'Orphaned',
        maintainerTooltip: {
            maintainedBy: 'Þis pakagge is beynge mayntened by {{ name }}',
            noMaintainer: 'Þis pakagge ne is beynge mayntened',
        },
        pagination: {
            previous: 'bak',
            next: 'nexte',
        },
    },
    packageDetails: {
        table: {
            name: 'Name',
            version: 'Version',
            maintainer: 'Mayntener',
            dependencies: 'Dependencies',
            requiredBy: 'Required By',
            lastUpdatedAt: 'Laste Updated At',
        },
        orphaned: 'Orphaned',
        noResults: 'None',
        openInGithub: 'Open In Github',
        view: 'Vewen',
        howToInstall: {
            title: 'How to Install',
            step1: 'Steppe 1: Setup Pacstall',
            step2: 'Steppe 2: Install {{ name }}',
        },
        dependenciesModal: {
            title: 'Dependencies',
            buildDependencies: 'Bylde Dependencies',
            optionalDependencies: 'Optional Dependencies',
            runtimeDependencies: 'Runtyme Dependencies',
            pacstallDependencies: 'Pacstall Dependencies',
            name: 'Name',
            close: 'Closen',
            provider: 'Provider',
            noDescription: 'Description ne is available',
            version: 'Version',
        },
        requiredByModal: {
            title: 'Required By',
            name: 'Name',
            provider: 'Provider',
            close: 'Closen',
            noDescription: 'Description ne is available',
        },
    },
}
