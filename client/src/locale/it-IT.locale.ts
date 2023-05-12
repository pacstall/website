import Locale from './locale'

export default <Locale>{
    home: {
        title: 'Pacstall',
        subtitle: "L'AUR per Ubuntu",
        cards: {
            whyDifferent: {
                title: 'Perché è diverso da qualunque altro gestore di pacchetti?',
                description:
                    'Pacstall usa la base stabile di Ubuntu ma ti permette di usare ' +
                    "software all'avanguardia senza compromessi, così non dovrai " +
                    'preoccuparti di patch di sicurezza o nuove funzionalità.',
            },
            howItWorks: {
                title: 'Bene, come funziona?',
                description:
                    ' Pacstall accetta file conosciuti come <0>pacscripts</0> (simili' +
                    ' a PKGBUILDs) che contengono i contenuti necessari per costruire i pacchetti,' +
                    ' e li trasforma in eseguibili sul tuo sistema.',
            },
        },
        installationInstructions: 'Istruzioni di installazione',
        showcase: {
            title: 'Vetrina',
            packageSearch: 'Ricerca di Pacchetti',
        },
    },
    navbar: {
        title: 'Pacstall',
        contribute: {
            title: 'Contribuisci',
            workOnFeatures: 'Lavora su nuove funzionalità',
            helpTranslate: 'Aiuta con le traduzioni',
            becomeAMaintainer: 'Diventa un mantenitore di pacchetti',
        },
        social: {
            title: 'Social Networks',
            discord: 'Discord',
            matrix: 'Matrix',
            reddit: 'Reddit',
            mastodon: 'Mastodon',
        },
        browse: {
            title: 'Cerca Pacchetti',
        },
        privacy: {
            title: 'Informativa sulla privacy',
        },
        install: 'Installa',
    },
    cookieConsent: {
        title: 'Avviso di cookie TL;DR',
        paragraphs: [
            'Ciao, sì, usiamo i cookie.',
            'Non ci piace darti informazioni fasulle o confusionarie. Usiamo i cookie solamente per funzionalità essenziali come le impostazioni del tema, localizzazione e autenticazione.',
            "Puoi leggere l'informativa sulla privacy completa  <0>qui <1/></0>.",
            "Continuando ad usare questo sito, <strong>acconsenti all'informativa sulla privacy</strong>.",
        ],
        accept: 'ok, va bene',
    },
    packageSearch: {
        dropdown: {
            package: 'Pacchetto',
            packageTooltip: 'Cerca nei nomi e nelle descrizioni dei pacchetti',
            maintainer: 'Mantenitore',
            maintainerTooltip: 'Cerca per nome ed email del mantenitore',
        },
        table: {
            name: 'Nome',
            maintainer: 'Mantenitore',
            version: 'Versione',
            install: 'Installa',
        },
        versionTooltip: {
            notInRegistry: 'Questo pacchetto non è nel registro di Repology',
            latest: "Questo pacchetto è aggiornato all'ultima versione",
            hasPatchUpdate: 'Questo pacchetto ha una patch disponibile',
            hasMinorUpdate:
                'Questo pacchetto ha un aggiornamento minore disponibile',
            hasMajorUpdate:
                'Questo pacchetto ha un aggiornamento importante disponibile',
            isGit: 'Questo pacchetto è un pacchetto Git',
        },
        noResults:
            'Non stai trovando quello che cerchi? <0>Crea una richiesta!</0>',
        search: 'Cerca',
        orphaned: 'Orfano',
        maintainerTooltip: {
            maintainedBy: 'Questo pacchetto è mantenuto da {{ name }}',
            noMaintainer: 'Questo pacchetto non è mantenuto da nessuno',
        },
        pagination: {
            previous: 'precedente',
            next: 'successivo',
        },
    },
    packageDetails: {
        table: {
            name: 'Nome',
            version: 'Versione',
            maintainer: 'Mantenitore',
            dependencies: 'Dipendenze',
            requiredBy: 'Richiesto da',
        },
        orphaned: 'Orfano',
        noResults: 'Nessuno',
        openInGithub: 'Apri su Github',
        view: 'Vedi',
        howToInstall: {
            title: 'Come Installare',
            step1: 'Passo 1: Installare Pacstall',
            step2: 'Passo 2: Installare {{ name }}',
        },
        dependenciesModal: {
            title: 'Dipendenze',
            buildDependencies: 'Costruisci Dipendenze',
            optionalDependencies: 'Dipendenze Facoltative',
            runtimeDependencies: 'Dipendenze Runtime',
            pacstallDependencies: 'Dipendenze Pacstall',
            name: 'Nome',
            close: 'Chiudi',
            provider: 'Fornitore',
            noDescription: 'Nessuna descrizione disponibile',
        },
        requiredByModal: {
            title: 'Richiesto da',
            name: 'Nome',
            provider: 'Fornitore',
            close: 'Chiudi',
            noDescription: 'Nessuna descrizione disponibile',
        },
    },
}
