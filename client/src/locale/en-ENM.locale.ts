import Locale from './locale'

export default <Locale>{
    home: {
        title: 'Pacstall',
        subtitle: 'Þe AUR for Ubuntu',
        cards: {
            whyDifferent: {
                title: 'Why is þis ani different þan any oþer pakagge manager?',
                description:
                    'Pacſtall noteþ þe ſtable baſe of Ubuntu but alloweþ þe to noteſt ' +
                    'bledynge egge ſofteware wiþ litel to no compromiſes, ſo þou ne doſt ' +
                    'have to troubleþ þeſelf wiþ ſecurite pacches or newe fetures.',
            },
            howItWorks: {
                title: 'How þenne doþ hit werk?',
                description:
                    ' Pacſtall takeþ in fyles witen as <0>pacſcripts</0> (like' +
                    ' PKGBUILDs) þat holden þe neceſsarye contents to bylden pakagges,' +
                    ' and byldeþ þem in-to executables on þi ſyſteme.',
            },
        },
        installationInstructions: 'Inſtallacioun Inſtrucciouns',
        showcase: {
            title: 'Schewecas',
            packageSearch: 'Pakagge Serchen',
        },
    },
    navbar: {
        title: 'Pacſtall',
        contribute: {
            title: 'Contribute',
            workOnFeatures: 'Werken on newe fetures',
            helpTranslate: 'Helpen wiþ tranſlaciouns',
            becomeAMaintainer: 'Becomen a pakagge mayntener',
        },
        social: {
            title: 'Social Nettwerks',
            discord: 'Diſcord',
            matrix: 'Matrix',
            reddit: 'Reddit',
            lemmy: 'Lemmy',
            mastodon: 'Maſtodon',
        },
        browse: {
            title: 'Browſen Pakagges',
        },
        privacy: {
            title: 'Privacie Policie',
        },
        install: 'Inſtallen',
    },
    cookieConsent: {
        title: 'Cookie notice TL;DR',
        paragraphs: [
            'Hail, we are notynge cookies on ure webſite.',
            'We ne doþ wiſshe to ȝiven ȝe miſledynge ne miſtakynge informacion. We oonly noten cookies for eſsential fetures ſwich as teme ſetynges, localiſacioun and auþentikacioun.',
            'Ȝe can rede þe ful privacie policie <0>her <1/></0>.',
            'By continuenge to noten þis ſite, ȝe <strong>ȝivest eower agrement to þe privacie policie</strong>.',
        ],
        accept: "God be wi' ye",
    },
    packageSearch: {
        dropdown: {
            package: 'Pakagge',
            packageTooltip: 'Serchen in pakagge names and deſcripciouns',
            maintainer: 'Mayntener',
            maintainerTooltip: 'Serchen by mayntener names and emailes',
        },
        table: {
            name: 'Name',
            maintainer: 'Mayntener',
            version: 'Verſion',
            install: 'Inſtall',
        },
        versionTooltip: {
            notInRegistry: 'Þis pakagge ne is in þe Repologgye reggiſtrye',
            latest: 'Þis pakagge is þe lateſt version',
            patch: 'Þis pakagge haveþ a patche update available',
            minor: 'Þis pakagge haveþ a minor update available',
            major: 'Þis pakagge haveþ a major update available',
            isGit: 'Þis pakagge is a Git pakagge',
        },
        noResults: 'Ne findynge what ȝe wiſsheſt? <0>Create a request!</0>',
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
            version: 'Verſion',
            maintainer: 'Mayntener',
            dependencies: 'Dependencies',
            requiredBy: 'Required By',
            lastUpdatedAt: 'Laſte Updated At',
        },
        orphaned: 'Orphaned',
        noResults: 'None',
        openInGithub: 'Open In Github',
        view: 'Vewen',
        howToInstall: {
            title: 'How to Inſtall',
            step1: 'Steppe 1: Setup Pacſtall',
            step2: 'Steppe 2: Inſtall {{ name }}',
        },
        dependenciesModal: {
            title: 'Dependencies',
            buildDependencies: 'Bylde Dependencies',
            optionalDependencies: 'Optional Dependencies',
            runtimeDependencies: 'Runtyme Dependencies',
            pacstallDependencies: 'Pacſtall Dependencies',
            name: 'Name',
            close: 'Closen',
            provider: 'Provider',
            noDescription: 'Deſcription ne is available',
            version: 'Verſion',
        },
        requiredByModal: {
            title: 'Required By',
            name: 'Name',
            provider: 'Provider',
            close: 'Cloſen',
            noDescription: 'Deſcription ne is available',
        },
    },
}
