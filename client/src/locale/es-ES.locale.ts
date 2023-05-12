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
                    'así que no tiene que preocuparse por los parches de seguridad ' +
                    'o las nuevas funciones.',
            },
            howItWorks: {
                title: '¿Pues, cómo funciona?',
                description:
                    ' Pacstall toma archivos conocidos como <0>pacscripts</0> (similares' +
                    ' a PKGBUILDs) que contienen lo que es necesario para construir paquetes, ' +
                    ' y los construye en ejecutables en su sistema.',
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
        title: 'Condiciones de las cookies TL;DR',
        paragraphs: [
            'Hola, sí, utilizamos cookies.',
            'No nos gusta darle información engañosa o confusa. Solo utilizamos cookies para funciones esenciales como la configuración de temas, la localización y la autenticación.',
            'Puede leer la política de privacidad completa <0>aquí <1/></0>',
            'Al continuar utilizando este sitio, usted <0>acepta esta política de privacidad</0>.',
        ],
        accept: 'ok, genial',
    },
    packageSearch: {
        dropdown: {
            package: 'Paquete',
            packageTooltip: 'Buscar nombres y descripciones de paquetes',
            maintainer: 'Mantenedor',
            maintainerTooltip:
                'Búsquedas por nombres y correos electrónicos de los mantenedores',
        },
        table: {
            name: 'Nombre',
            maintainer: 'Mantenedor',
            version: 'Versión',
            install: 'Instalar',
        },
        versionTooltip: {
            notInRegistry: 'Este paquete no está en el registro de Repology',
            latest: 'Este paquete es la última versión',
            patch:
                'Este paquete dispone de un parche de actualización',
            minor: 'Este paquete dispone de una actualización menor',
            major: 'Este paquete dispone de una actualización mejor',
            isGit: 'Este paquete es un paquete Git',
        },
        orphaned: 'Huérfanos',
        noResults: '¿No encuentra lo que busca? <0>¡Crear una solicitud!</0>',
        search: 'Buscar',
        maintainerTooltip: {
            maintainedBy:
                'El mantenimiento de este paquete corre a cargo de {{ name }}',
            noMaintainer: 'Este paquete no se mantiene',
        },
        pagination: {
            previous: 'anterior',
            next: 'próximo',
        },
    },
    packageDetails: {
        table: {
            name: 'Nombre',
            version: 'Versión',
            maintainer: 'Mantenedor',
            dependencies: 'Dependencias',
            requiredBy: 'Requierdo por',
        },
        orphaned: 'Huérfanos',
        noResults: 'Nada',
        openInGithub: 'Abrir en Github',
        view: 'Ver',
        howToInstall: {
            title: 'Cómo Instalar',
            step1: 'Paso 1: Instalar Pacstall',
            step2: 'Paso 2: Instalar {{ name }}',
        },
        dependenciesModal: {
            title: 'Dependencias',
            buildDependencies: 'Dependencias de Construir',
            optionalDependencies: 'Dependencias Opcionales',
            runtimeDependencies: 'Dependencias del tiempo de Ejecución',
            pacstallDependencies: 'Dependencias de Pacstall',
            name: 'Nombre',
            close: 'Cerrar',
            provider: 'Proveedor',
            noDescription: 'No hay descripción disponible',
        },
        requiredByModal: {
            title: 'Requerido por',
            name: 'Nombre',
            provider: 'Proveedor',
            close: 'Cerrar',
            noDescription: 'No hay descripción disponible',
        },
    },
}
