import Locale from './locale'

export default <Locale>{
    home: {
        title: 'Pacstall',
        subtitle: 'El AUR para Ubuntu',
        cards: {
            whyDifferent: {
                title: '¿Por qué es diferente a cualquier otro administrador de paquetes?',
                description:
                    'Pacstall usa la base estable de Ubuntu pero le permite ' +
                    'usar software de última generación con poco o ningún compromiso, ' +
                    'por lo que no tiene que preocuparse por los parches de seguridad ' +
                    'o las nuevas funciones.',
            },
            howItWorks: {
                title: '¿Cómo funciona entonces?',
                description:
                    ' Pacstall takes in files known as <link>pacscripts</link> (similar' +
                    ' to PKGBUILDs) that contain the necessary contents to build packages,' +
                    ' and builds them into executables on your system.',
            },
        },
        installationInstructions: 'Instrucciones de Instalación',
        showcase: {
            title: 'Escaparate',
            packageSearch: 'Búsqueda de Paquetes',
        },
    },
    navbar: {
        title: 'Pacstall',
        contribute: {
            title: 'Contribuir',
            workOnFeatures: 'Trabaja en nuevas características',
            helpTranslate: 'Ayuda con las traducciones',
            becomeAMaintainer: 'Conviértase en un mantenedor de paquetes',
        },
        social: {
            title: 'Redes Sociales',
            discord: 'Discord',
            matrix: 'Matrix',
            reddit: 'Reddit',
            mastodon: 'Mastodon',
        },
        browse: {
            title: 'Explorar Paquetes',
        },
        privacy: {
            title: 'La Política de Privacidad',
        },
        install: 'Instalar',
    },
    cookieConsent: {
        title: 'Cookie notice TL;DR',
        paragraphs: [
            'Hi, yes, we use cookies.',
            "We don't like to give you misleading or confusing information. We only use cookies for essential features such as theme settings, localization and authentication.",
            'You can read the full privacy policy <0>here <1/></0>',
            'By continuing to use this site, you <strong>give your agreement to the privacy policy</strong>.',
        ],
        accept: 'ok, nice',
    },
}
