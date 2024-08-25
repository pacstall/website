import Locale from './locale'

export default <Locale>{
    home: {
        title: 'Pæxtal',
        subtitle: 'Þone AUR for Ubunte',
        cards: {
            whyDifferent: {
                title: 'Hƿȳ is þis ǣniġ unġelīċ þonne ǣniġ ōþer pæcca āganere?',
                description:
                    'Þone strangne flōr Pæxtal brȳcþ Ubunte būtan þē ālīefest brūcan ' +
                    'blēdende eċġe sōfteƿaru ƿiþ lȳtel oþþe nāhƿæþer dūnesīdan, sƿā þē ne ' +
                    'fremest tō dreċest þīnes silfes onbūtan felþu plæċċa āhƿæþer nīeƿe maganþus.',
            },
            howItWorks: {
                title: 'Hū hit ƿyrċeþ þonne?',
                description:
                    ' Pæxtal tæcþ filas ġecnāƿen eallsƿā <0>pacscripts</0> (līciaþ' +
                    ' PKGBUILDas) þæt healdaþ sēo nīedan bitan tō bytlan pæccas,' +
                    ' and bytleþ hīe rinnendlīċ intō on þīn system.',
            },
        },
        installationInstructions: 'Onstellinga Tǣċinga',
        showcase: {
            title: 'Sċēaƿbox',
            packageSearch: 'Pæcca Sēċan',
        },
    },
    navbar: {
        title: 'Pæxtal',
        contribute: {
            title: 'Helpan',
            workOnFeatures: 'Ƿyrċan on nieƿe maganþus',
            helpTranslate: 'Helpan ƿiþ ƿendingum',
            becomeAMaintainer: 'Becuman pæcca cēpere',
        },
        social: {
            title: 'Seg Nettƿeorc',
            discord: 'Discord',
            matrix: 'Matrix',
            reddit: 'Reddit',
            lemmy: 'Lemmy',
            mastodon: 'Mastodon',
        },
        browse: {
            title: 'Lōcian Pæccas',
        },
        privacy: {
            title: 'Gehȳdendenes Lagu',
        },
        install: 'Onstellan',
    },
    cookieConsent: {
        title: 'Kaka hēdan TL;NR',
        paragraphs: [
            'Ƿes hāl, ƿē sind brūcinge kakan on ūre ƿebbplæċe.',
            'Ƿē ne līciaþ tō ġiefan þē mislǣdung ōþer mistacing ƿord. Ānlīċ brūcaþ ƿē kakan for nīedende maganþus sƿā sƿā anyettanness settunga, ƿending and cēpung.',
            '<0>Hēr<1/></0> ƿāst þū rǣtst þone full gehȳdendenes lagu.',
            'Bī þurhƿunung tō brūcan þisse ƿebbplæċe, <strong>ġiefst ēoƿ ēoƿer ġieldscype tō þisse gehȳdendenes lagu</strong>.',
        ],
        accept: 'Sōþlīċe, hit īs gōd',
    },
    packageSearch: {
        dropdown: {
            package: 'Pæcca',
            packageTooltip: 'Sēċan in pæcca naman and besċrīfunga',
            maintainer: 'Pæcca Cēpere',
            maintainerTooltip: 'Hit Sēċeþ bī pæcca cēpere naman and bōcstafas',
        },
        table: {
            name: 'Nama',
            maintainer: 'Pæcca Cēpere',
            version: 'Weorþ',
            install: 'Onstellan',
        },
        versionTooltip: {
            notInRegistry: 'Þisse pæcca ne is in sēo Repology bōcsċielfe',
            latest: 'Þisse pæcca is þæt latost ƿeorþ',
            patch: 'Þisse pæcca hæfþ plæċċ uptīma forfēġbǣre',
            minor: 'Þisse pæcca hæfþ minor uptīma forfēġbǣre',
            major: 'Þisse pæcca hæfþ major uptīma forfēġbǣre',
            isGit: 'Þisse pæcca is Git pæcca',
        },
        noResults:
            'Ne finding þæt hƿelċ þū earon lōcung for, eart þu? <0>Ġesċiepest ġesċeaft!</0>',
        search: 'Sēċan',
        orphaned: 'Unierfa',
        maintainerTooltip: {
            maintainedBy: 'Þisse pæcca is cēpende bī {{ name }}',
            noMaintainer: 'Þisse pæcca ne is cēpende',
        },
        pagination: {
            previous: 'bæc',
            next: 'nīehst',
        },
    },
    packageDetails: {
        table: {
            name: 'Nama',
            version: 'Weorþ',
            maintainer: 'Pæcca Cēpere',
            dependencies: 'Æfhōniġhādas',
            requiredBy: 'Nīedende bī',
            lastUpdatedAt: 'Latost Uptīmiaþ Æt',
        },
        orphaned: 'Unierfa',
        noResults: 'Nān',
        openInGithub: 'Openian In Github',
        view: 'Lōcian',
        howToInstall: {
            title: 'Hū tō Onstellan',
            step1: 'Stæpe I: Upsettan Pæxtal',
            step2: 'Stæpe II: Onstellan {{ name }}',
        },
        dependenciesModal: {
            title: 'Æfhōniġhādas',
            buildDependencies: 'Æfhōniġhādas Byldan',
            optionalDependencies: 'Æfhōniġhādas Ġecoren',
            runtimeDependencies: 'Æfhōniġhādas Rinntīma',
            pacstallDependencies: 'Æfhōniġhāda Pæxtala',
            name: 'Nama',
            close: 'Clȳsan',
            provider: 'Profidere',
            noDescription: 'Nān ġereċċendeness gearo beon',
            version: 'Weorþ',
        },
        requiredByModal: {
            title: 'Nīedende bī',
            name: 'Nama',
            provider: 'Profidere',
            close: 'Clȳsan',
            noDescription: 'Nān ġereċċendeness gearo beon',
        },
    },
}
