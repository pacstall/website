import Locale from './locale'

export default <Locale>{
    home: {
        title: 'Pacstall',
        subtitle: 'Ett AUR för Ubuntu',
        cards: {
            whyDifferent: {
                title: 'Vad är skillnaden mellan detta och alla andra pakethanterare?',
                description:
                    'Pacstall använder den stabilla grunden av Ubuntu men tillåter ' +
                    'dig att använda den senaste mjukvaran med få till inga kompromisser, så du behöver ' +
                    'inte oroa dig om att stå utan säkerhetspatcher eller nya funktioner.',
            },
            howItWorks: {
                title: 'Hur funkar det då?',
                description:
                    ' Pacstall använder filer som kallas <0>pacscripts</0> (liknar' +
                    ' PKGBUILDs) som har instruktioner för att bygga mjukvarupaket' +
                    ' och sedan installera dem som exekverbara filer på ditt system.',
            },
        },
        installationInstructions: 'Installationsguide',
        showcase: {
            title: 'Demonstration',
            packageSearch: 'Sök efter paket',
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
            'Hej, ja, vi använder cookies.',
            'Vi tänker inte ge dig felaktig eller förvirrande information. Vi använder bara cookies för nödvändiga funktioner t.ex. temainställningar, lokalisering, och autentisering.',
            'Du kan läsa hela integritetspolicyn <0>här <1/></0>.',
            'Genom att fortsätta använda vår hemsida så <strong>ger du ditt samtycke till integritetspolicyn</strong>.',
        ],
        accept: 'ok, coolt',
    },
    packageSearch: {
        dropdown: {
            package: 'Paket',
            packageTooltip: 'Söker efter paketnamn och beskrivningar',
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
            notInRegistry: 'Detta paket finns inte i Repologys databas',
            latest: 'Det här paketet är den senaste versionen',
            patch: 'Det här paketet har en patchuppdatering tillgänglig',
            minor: 'Det här paketet har en liten uppdatering tillgänglig',
            major: 'Det här paketet har en stor uppdatering tillgänglig',
            isGit: 'Det här paketet är ett Git paket',
        },
        noResults:
            'Hittar du inte vad du letar efter? <0>Skicka en förfrågan!</0>',
        search: 'Sök',
        orphaned: 'Ej underhållen',
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
            dependencies: 'Beror på',
            requiredBy: 'Krävs av',
            lastUpdatedAt: 'Senast uppdaterad',
        },
        orphaned: 'Ej underhållen',
        noResults: 'Inga resultat',
        openInGithub: 'Öppna i Github',
        view: 'Visa',
        howToInstall: {
            title: 'Installationsguide',
            step1: 'Steg 1: Installera Pacstall',
            step2: 'Steg 2: Installera {{ name }}',
        },
        dependenciesModal: {
            title: 'Beror på',
            buildDependencies: 'Byggberoenden',
            optionalDependencies: 'Valfria beroenden',
            runtimeDependencies: 'Exekveringsberoenden',
            pacstallDependencies: 'Pacstall beroenden',
            name: 'Namn',
            close: 'Stäng',
            provider: 'Källa',
            noDescription: 'Ingen beskrivning finns',
            version: 'Version',
        },
        requiredByModal: {
            title: 'Krävs av',
            name: 'Namn',
            provider: 'Källa',
            close: 'Stäng',
            noDescription: 'Ingen beskrivning finns',
        },
    },
}
