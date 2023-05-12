import Locale from './locale'

export default <Locale>{
    home: {
        title: 'Pacstall',
        subtitle: 'Das AUR für Ubuntu',
        cards: {
            whyDifferent: {
                title: 'Was unterscheidet es von anderen Paketmanagern?',
                description:
                    'Pacstall verwendet die stabile Basis von Ubuntu, erlaubt aber die Nutzung ' +
                    'von modernster Software ohne oder mit nur geringen Kompromissen, sodass du ' +
                    'dich nicht um Sicherheitspatches oder neue Funktionen kümmern musst.',
            },
            howItWorks: {
                title: 'Wie funktioniert es?',
                description:
                    ' Pacstall nimmt Dateien namens <0>pacscripts</0> (ähnlich wie PKGBUILDs),' +
                    ' die die notwendigen Inhalte zum Erstellen von Paketen enthalten, und erstellt' +
                    ' daraus ausführbare Dateien auf deinem System.',
            },
        },
        installationInstructions: 'Installationsanleitung',
        showcase: {
            title: 'Ausstellung',
            packageSearch: 'Paket suchen',
        },
    },
    navbar: {
        title: 'Pacstall',
        contribute: {
            title: 'Mitwirken',
            workOnFeatures: 'An neuen Funktionen arbeiten',
            helpTranslate: 'Bei Übersetzungen helfen',
            becomeAMaintainer: 'Paketbetreuer werden',
        },
        social: {
            title: 'Soziale Netzwerke',
            discord: 'Discord',
            matrix: 'Matrix',
            reddit: 'Reddit',
            mastodon: 'Mastodon',
        },
        browse: {
            title: 'Pakete durchsuchen',
        },
        privacy: {
            title: 'Datenschutz-Bestimmungen',
        },
        install: 'Installieren',
    },
    cookieConsent: {
        title: 'Cookie-Hinweis TL;DR',
        paragraphs: [
            'Hallo, ja, wir verwenden Cookies.',
            'Wir möchten dir keine irreführenden oder verwirrenden Informationen geben. Wir verwenden nur Cookies für wesentliche Funktionen wie Theme Einstellungen, Lokalisierung und Authentifizierung.',
            'Du kannst die vollständigen Datenschutzbestimmungen <0>hier <1/></0> lesen.',
            'Durch die weitere Nutzung dieser Website <strong>stimmst du den Datenschutzbestimmungen zu</strong>.',
        ],
        accept: 'Ok, gut',
    },
    packageSearch: {
        dropdown: {
            package: 'Paket',
            packageTooltip: 'Sucht in Paketnamen und Beschreibungen',
            maintainer: 'Betrauer',
            maintainerTooltip: 'Sucht nach Namen und E-Mails von Betreuern',
        },
        table: {
            name: 'Name',
            maintainer: 'Betreuer',
            version: 'Version',
            install: 'Installieren',
        },
        versionTooltip: {
            notInRegistry: 'Dieses Paket ist nicht im Repology-Register',
            latest: 'Diese Version ist die aktuellste',
            patch: 'Es gibt ein Patch-Update für dieses Paket',
            major: 'Es gibt ein Minor-Update für dieses Paket',
            minor: 'Es gibt ein Major-Update für dieses Paket',
            isGit: 'Dieses Paket ist ein Git-Paket',
        },
        noResults: 'Findest du nicht, wonach du suchst? <0>Erstelle eine Anfrage!</0>',
        search: 'Suche',
        orphaned: 'Verwaist',
        maintainerTooltip: {
            maintainedBy: 'Dieses Paket wird von {{ name }} betreut',
            noMaintainer: 'Dieses Paket wird nicht betreut',
        },
        pagination: {
            previous: 'zurück',
            next: 'weiter',
        },
    },
    packageDetails: {
        table: {
            name: 'Name',
            version: 'Version',
            maintainer: 'Betreuer',
            dependencies: 'Abhängigkeiten',
            requiredBy: 'Wird benötigt von',
        },
        orphaned: 'Verwaist',
        noResults: 'Keine',
        openInGithub: 'In Github öffnen',
        view: 'Ansehen',
        howToInstall: {
            title: 'Installation',
            step1: 'Schritt 1: Pacstall einrichten',
            step2: 'Schritt 2: {{ name }} installieren',
        },
        dependenciesModal: {
            title: 'Abhängigkeiten',
            buildDependencies: 'Build-Abhängigkeiten',
            optionalDependencies: 'Optionale Abhängigkeiten',
            runtimeDependencies: 'Laufzeit-Abhängigkeiten',
            pacstallDependencies: 'Pacstall-Abhängigkeiten',
            name: 'Name',
            close: 'Schließen',
            provider: 'Anbieter',
            noDescription: 'Keine Beschreibung verfügbar',
        },
        requiredByModal: {
            title: 'Wird benötigt von',
            name: 'Name',
            provider: 'Anbieter',
            close: 'Schließen',
            noDescription: 'Keine Beschreibung verfügbar',
        },
    },
}
