import Locale from './locale'

export default <Locale>{
    home: {
        title: 'Pacstall',
        subtitle: 'Alternativa AUR pentru Ubuntu',
        cards: {
            whyDifferent: {
                title: 'In fel ce este diferit de alte package managere?',
                description:
                    'Pacstall folosește bază stabilă a Ubuntu dar iți permite să folosești' +
                    ' software bleeding edge cu puține compromisuri, așa că nu trebuie să te' +
                    ' îngrijorezi de patch-uri de securitate sau de noi funcționalități.',
            },
            howItWorks: {
                title: 'Cum funcționează?',
                description:
                    'Pacstall ia fișiere cunoscute ca <0>pacscripts</0> (similar' +
                    ' cu PKGBUILD) care conțin conținutul necesar pentru a crea pachete instalabile,' +
                    ' și le instalează în sistemul tău.',
            },
        },
        installationInstructions: 'Instrucțiuni de instalare',
        showcase: {
            title: 'Prezentare',
            packageSearch: 'Căutare de aplicații',
        },
    },
    navbar: {
        title: 'Pacstall',
        contribute: {
            title: 'Contribuie',
            workOnFeatures: 'Lucrează la funcționalitățiile noi',
            helpTranslate: 'Ajută cu traducerile',
            becomeAMaintainer: 'Adaugă aplicații noi',
        },
        social: {
            title: 'Rețele sociale',
            discord: 'Discord',
            matrix: 'Matrix',
            reddit: 'Reddit',
            mastodon: 'Mastodon',
        },
        browse: {
            title: 'Caută aplicații',
        },
        privacy: {
            title: 'Politica de confidențialitate',
        },
        install: 'Instalează',
    },
    cookieConsent: {
        title: 'Politica de confidențialitate, pe scurt',
        paragraphs: [
            'Salut, da, folosim cookies.',
            'Nu vrem să îți oferim informații greșite sau confuze. Folosim cookies doar pentru funcționalități esențiale, cum ar fi setările de temă, localizarea și autentificarea.',
            'Poți citi politica de confidențialitate completă <0>aici <1/></0>.',
            'Continuând să folosești acest site, <0>îți dai acordul pentru politica de confidențialitate</0>.',
        ],
        accept: 'am înțeles',
    },
    packageSearch: {
        dropdown: {
            package: 'Aplicații',
            packageTooltip: 'Caută aplicații după nume sau descriere',
            maintainer: 'Menținători',
            maintainerTooltip: 'Caută aplicații după menținatori sau emailuri',
        },
        table: {
            name: 'Nume',
            maintainer: 'Menținut de',
            version: 'Versiune',
            install: 'Instalare',
        },
        versionTooltip: {
            notInRegistry:
                'Această aplicație nu este înregistrată în registrul Repology',
            latest: 'Această aplicație este la ultima versiune',
            hasPatchUpdate:
                'Această aplicație are o actualizare de securitate disponibilă',
            hasMinorUpdate:
                'Această aplicație are o actualizare minoră disponibilă',
            hasMajorUpdate:
                'Această aplicație are o actualizare majoră disponibilă',
            isGit: 'Această aplicație este de tip Git',
        },
        orphaned: 'Nemenținut',
        noResults: 'Nu ai găsit ce căutai? <0>Scrie o cerere!</0>',
        search: 'Caută',
        maintainerTooltip: {
            maintainedBy: 'Această aplicație este întreținută de {{ name }}',
            noMaintainer: 'Această aplicație nu este întreținută',
        },
        pagination: {
            previous: 'înapoi',
            next: 'înainte',
        },
    },
    packageDetails: {
        table: {
            name: 'Nume',
            version: 'Versiune',
            maintainer: 'Menținut de',
            dependencies: 'Dependențe',
            requiredBy: 'Necesar pentru',
        },
        orphaned: 'Nemenținut',
        noResults: 'Fără',
        openInGithub: 'Deschide în Github',
        view: 'Vezi',
        howToInstall: {
            title: 'Cum se instalează',
            step1: 'Pasul 1: Instalează Pacstall',
            step2: 'Pasul 2: Instalează {{ name }}',
        },
        dependenciesModal: {
            title: 'Dependențe',
            buildDependencies: 'Dependențe de construcție',
            optionalDependencies: 'Dependențe opționale',
            runtimeDependencies: 'Dependențe de rulare',
            pacstallDependencies: 'Dependențe Pacstall',
            name: 'Nume',
            close: 'Închide',
            provider: 'Furnizor',
            noDescription: 'Nu există o descriere disponibilă',
        },
        requiredByModal: {
            title: 'Necesar pentru',
            name: 'Nume',
            provider: 'Furnizor',
            close: 'Închide',
            noDescription: 'Nu există o descriere disponibilă',
        },
    },
}
