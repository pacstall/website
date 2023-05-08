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
}
