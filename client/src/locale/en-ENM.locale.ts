import Locale from './locale'

export default <Locale>{
    home: {
        title: 'Pacſtall',
        subtitle: 'Þe AUR foꝛ Ubuntu',
        cards: {
            whyDifferent: {
                title: 'Why is þis ani different þan any oþer pakagge manager?',
                description:
                    'Pacſtall noteþ þͤ ſtable baſe of Ubuntu but alloweþ þͤ to noteſt ' +
                    'bledynge egge ſofteware wiþ litel to no compꝛomiſes, ſo þͧ ne doſt ' +
                    'haue to troubleþ þeſelf wiþ ſecurite pacches oꝛ newe fetures⸵',
            },
            howItWorks: {
                title: 'How þenne doþ hit werk?',
                description:
                    ' Pacſtall takeþ in fyles witen as <0>pacſcripts</0> (like' +
                    ' PKGBUILDs) þat holden þͤ neceſsarye contents to bylden pakagges‧' +
                    ' ⁊ byldeþ þem in-to executables on þi ſyſteme⸵',
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
            discord: 'Diſcoꝛd',
            matrix: 'Matrix',
            reddit: 'Reddit',
            lemmy: 'Lemmy',
            mastodon: 'Maſtodon',
        },
        browse: {
            title: 'Browſen Pakagges',
        },
        privacy: {
            title: 'Priuacie Policie',
        },
        install: 'Inſtallen',
    },
    cookieConsent: {
        title: 'Cookie notice TL;DR',
        paragraphs: [
            'Hail‧ we are notynge cookies on ure webſite⸵',
            'We ne doþ wiſshe to ȝiuen ȝe miſledynge ne miſtakynge infoꝛmacion⸵ We oonly noten cookies foꝛ eſsential fetures ſwich as teme ſetynges, localiſacioun ⁊ auþentikacioun⸵',
            'Ȝe can rede þͤ ful pꝛiuacie policie <0>her <1/></0>⸵',
            'By continuenge to noten þis ſite‧ ȝe <strong>ȝiuest eower agrement to þͤ pꝛiuacie policie</strong>⸵',
        ],
        accept: "God be wi' yͤ",
    },
    packageSearch: {
        dropdown: {
            package: 'Pakagge',
            packageTooltip: 'Serchen in pakagge names ⁊ deſcripciouns',
            maintainer: 'Mayntener',
            maintainerTooltip: 'Serchen by mayntener names ⁊ emailes',
        },
        table: {
            name: 'Name',
            maintainer: 'Mayntener',
            version: 'Verſion',
            install: 'Inſtall',
        },
        versionTooltip: {
            notInRegistry: 'Þis pakagge ne is in þͤ Repologgye reggiſtrye',
            latest: 'Þis pakagge is þͤ lateſt version',
            patch: 'Þis pakagge haueþ a patche update auailable',
            minor: 'Þis pakagge haueþ a minoꝛ update auailable',
            major: 'Þis pakagge haueþ a majoꝛ update auailable',
            isGit: 'Þis pakagge is a Git pakagge',
        },
        noResults: 'Ne findynge what ȝe wiſsheſt? <0>Create a requeſt!</0>',
        search: 'Serchen',
        orphaned: 'Oꝛphaned',
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
            step1: 'Steppe I: Setup Pacſtall',
            step2: 'Steppe II: Inſtall {{ name }}',
        },
        dependenciesModal: {
            title: 'Dependencies',
            buildDependencies: 'Bylde Dependencies',
            optionalDependencies: 'Optional Dependencies',
            runtimeDependencies: 'Runtyme Dependencies',
            pacstallDependencies: 'Pacſtall Dependencies',
            name: 'Name',
            close: 'Cloſen',
            provider: 'Prouider',
            noDescription: 'Deſcription ne is auailable',
            version: 'Verſion',
        },
        requiredByModal: {
            title: 'Required By',
            name: 'Name',
            provider: 'Prouider',
            close: 'Cloſen',
            noDescription: 'Deſcription ne is auailable',
        },
    },
}
