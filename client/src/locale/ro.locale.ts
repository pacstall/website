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
                    'Pacstall ia fișiere cunoscute ca <link>pacscripts</link> (similar' +
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
            'Continuând să folosești acest site, <strong>îți dai acordul pentru politica de confidențialitate</strong>.',
        ],
        accept: 'am înțeles',
    },
}
