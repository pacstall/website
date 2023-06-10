import Locale from './locale'

export default <Locale>{
    home: {
        title: 'Pacstall',
        subtitle: 'Ett AUR för Ubuntu',
        cards: {
            whyDifferent: {
                title: 'Vad är skillnaden mellan detta och alla andra pakethanterare?',
                description:
                    'Pacstall andvänder den stabilla grunden av Ubuntu men tillåter ' +
                    'dig att andvända den senaste mjukvaran utan orsaker, så du behöver' +
                    'inte oroa dig för säkerhetspatcher eller nya funktioner.',
            },
            howItWorks: {
                title: 'Hur funkar det då?',
                description:
                    ' Pacstall andvänder filer som kallas <0>pacscripts</0> (liknande' +
                    ' PKGBUILDs) som har informationen som krävs för att bygga paketen,' +
                    ' och bygger dem som körbara filer på ditt system.',
            },
        },
        installationInstructions: 'Installation Instruktioner',
        showcase: {
            title: 'Demonstration',
            packageSearch: 'Sök för paket',
        },
    },
    navbar: {
        title: 'Pacstall',
        contribute: {
            title: 'Bidra',
            workOnFeatures: 'Skapa nya funktioner',
            helpTranslate: 'Hjälp översätta Pacstall',
            becomeAMaintainer: 'Bli paketunderhållare',
        },
        social: {
            title: 'Sociala Nätverk',
            discord: 'Discord',
            matrix: 'Matrix',
            reddit: 'Reddit',
            lemmy: 'Lemmy',
            mastodon: 'Mastodon',
        },
        browse: {
            title: 'Bläddra Paket',
        },
        privacy: {
            title: 'Integritetspolicy',
        },
        install: 'Installera',
    },
    cookieConsent: {
        title: 'Cookiemeddelande TL;DR',
        paragraphs: [
            'Hej, ja, vi andvänder cookies.',
            'Vi tänker inte ge dig felaktig eller förvirrande information. Vi andvänder bara cookies för nödvändiga funktioner t.ex. temainställningar, lokalisering, och autentisering.',
            'Du kan läsa hela integritetspolicyn <0>här <1/></0>.',
            'Med att fortsätta andvända vår site så <strong>ger du ditt samtycke till integritetspolicyn</strong>.',
        ],
        accept: 'ok, coolt',
    },
    packageSearch: {
        dropdown: {
            package: 'Paket',
            packageTooltip: 'Söker i paketnamn och beskrivningar',
            maintainer: 'Underhållare',
            maintainerTooltip: 'Söker efter underhållares namn och email',
        },
        table: {
            name: 'Namn',
            maintainer: 'Underhållare',
            version: 'Version',
            install: 'Installera',
        },
        versionTooltip: {
            notInRegistry: 'Detta paket finns inte i Repology registret',
            latest: 'Detta paket är den senaste versionen',
            patch: 'Detta paket har en patchuppdatering tillgänglig',
            minor: 'Detta paket har en liten uppdatering tillgänglig',
            major: 'Detta paket har en stor uppdatering tillgänglig',
            isGit: 'Detta paket är ett Git paket',
        },
        noResults:
            'Hittar du inte vad du letar efter? <0>Skicka en förfrågan!</0>',
        search: 'Sök',
        orphaned: 'Oandvänd',
        maintainerTooltip: {
            maintainedBy: 'Detta paket underhålls av {{ name }}',
            noMaintainer: 'Detta paket har ingen underhållare',
        },
        pagination: {
            previous: 'förra',
            next: 'nästa',
        },
    },
    packageDetails: {
        table: {
            name: 'Namn',
            version: 'Version',
            maintainer: 'Underhållare',
            dependencies: 'Beroenden',
            requiredBy: 'Krävs av',
        },
        orphaned: 'Oandvänd',
        noResults: 'Inga resultat',
        openInGithub: 'Öpnna i Github',
        view: 'Visa',
        howToInstall: {
            title: 'Installationsguide',
            step1: 'Steg 1: Installera Pacstall',
            step2: 'Steg 2: Installera {{ name }}',
        },
        dependenciesModal: {
            title: 'Beroenden',
            buildDependencies: 'Bygg beroenden',
            optionalDependencies: 'Valfria beroenden',
            runtimeDependencies: 'Körtidsberoenden',
            pacstallDependencies: 'Pacstall beronden',
            name: 'Namn',
            close: 'Stäng',
            provider: 'Leverantör',
            noDescription: 'Ingen beskrivning finns',
        },
        requiredByModal: {
            title: 'Krävs av',
            name: 'Namn',
            provider: 'Leverantör',
            close: 'Stäng',
            noDescription: 'Ingen beskrivning finns',
        },
    },
}
