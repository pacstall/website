import Locale from './locale'

export default <Locale>{
    home: {
        title: 'Pacstall',
        subtitle: 'De AUR voor Ubuntu',
        cards: {
            whyDifferent: {
                title: 'Waarom is dit anders dan andere pakketbeheerders?',
                description:
                    'Pacstall gebruikt de stabiele basis van Ubuntu maar laat je ' +
                    'de nieuwste software gebruiken met vrijwel geen compromissen, zodat je je ' +
                    'geen zorgen hoeft te maken over beveiligingsupdates of nieuwe functies.',
            },
            howItWorks: {
                title: 'Maar hoe werkt het dan?',
                description:
                    ' Pacstall gebruikt bestanden genaamd <0>pacscripts</0> (vergelijkbaar' +
                    ' met PKGBUILDs) die alles bevatten om pakketten te maken,' +
                    ' en bouwt daarmee uitvoerbare bestanden op je systeem.',
            },
        },
        installationInstructions: 'Installatie Instructies',
        showcase: {
            title: 'Showcase',
            packageSearch: 'Pakketten Zoeken',
        },
    },
    navbar: {
        title: 'Pacstall',
        contribute: {
            title: 'Bijdragen',
            workOnFeatures: 'Werk mee aan nieuwe functies',
            helpTranslate: 'Help met vertalen',
            becomeAMaintainer: 'Word pakketonderhouder',
        },
        social: {
            title: 'Sociale Netwerken',
            discord: 'Discord',
            matrix: 'Matrix',
            reddit: 'Reddit',
            mastodon: 'Mastodon',
        },
        browse: {
            title: 'Pakketten Bekijken',
        },
        privacy: {
            title: 'Privacyverklaring',
        },
        install: 'Installeer',
    },
    cookieConsent: {
        title: 'Cookiemelding in het kort',
        paragraphs: [
            'Hoi, ja, wij gebruiken cookies.',
            'We geven je niet graag misleidende of verwarrende informatie. We gebruiken cookies alleen voor essentiële functies zoals themainstellingen, localisatie and authenticatie.',
            'Je kan de volledige privacyverklaring <0>hier <1/></0> lezen.',
            'Door deze site te blijven gebruiken, <strong>stem je in met de privacyverklaring</strong>.',
        ],
        accept: 'ok, leuk',
    },
    packageSearch: {
        dropdown: {
            package: 'Pakket',
            packageTooltip: 'Zoekt in pakketnamen en -omschrijvingen',
            maintainer: 'Onderhouder',
            maintainerTooltip: 'Zoekt op onderhoudernamen en -emails',
        },
        table: {
            name: 'Naam',
            maintainer: 'Onderhouder',
            version: 'Versie',
            install: 'Installeer',
        },
        versionTooltip: {
            notInRegistry: 'Dit pakket is niet beschikbaar in het Repology-archief',
            latest: 'Dit pakket heeft de laatste versie',
            patch: 'Voor dit pakket is een patch update beschikbaar',
            minor: 'Voor dit pakket is een minor update beschikbaar',
            major: 'Voor dit pakket is een major update beschikbaar',
            isGit: 'Dit pakket is een Git-pakket',
        },
        noResults: 'Niet gevonden wat je zocht? <0>Dien een verzoek in!</0>',
        search: 'Zoeken',
        orphaned: 'Verweesd',
        maintainerTooltip: {
            maintainedBy: 'Dit pakket wordt onderhouden door {{ name }}',
            noMaintainer: 'Dit pakket wordt niet onderhouden',
        },
        pagination: {
            previous: 'vorige',
            next: 'volgende',
        },
    },
    packageDetails: {
        table: {
            name: 'Naam',
            version: 'Versie',
            maintainer: 'Onderhouder',
            dependencies: 'Afhankelijkheden',
            requiredBy: 'Vereist Door',
        },
        orphaned: 'Verweesd',
        noResults: 'Geen',
        openInGithub: 'Open Op Github',
        view: 'Bekijk',
        howToInstall: {
            title: 'Installatie Instructies',
            step1: 'Stap 1: Pacstall Installeren',
            step2: 'Stap 2: Installeer {{ name }}',
        },
        dependenciesModal: {
            title: 'Afhankelijkheden',
            buildDependencies: 'Bouwafhankelijkheden',
            optionalDependencies: 'Optionele Afhankelijkheden',
            runtimeDependencies: 'Runtime Afhankelijkheden',
            pacstallDependencies: 'Pacstall Afhankelijkheden',
            name: 'Naam',
            close: 'Sluiten',
            provider: 'Aanbieder',
            noDescription: 'Geen omschrijving beschikbaar',
        },
        requiredByModal: {
            title: 'Vereist Door',
            name: 'Naam',
            provider: 'Aanbieder',
            close: 'Sluiten',
            noDescription: 'Geen omschrijving beschikbaar',
        },
    },
}
