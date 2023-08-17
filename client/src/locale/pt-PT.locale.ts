import Locale from './locale'

export default <Locale>{
    home: {
        title: 'Pacstall',
        subtitle: 'O AUR para Ubuntu',
        cards: {
            whyDifferent: {
                title: 'Por que é que isto é diferente de qualquer outro gestor de pacotes?',
                description:
                    'O Pacstall usa a base estável do Ubuntu mas permite-lhe usar o ' +
                    'software de última geração com poucos ou nenhuns compromissos, pelo que ' +
                    'não tem de se preocupar com correções de segurança ou novas funcionalidades.',
            },
            howItWorks: {
                title: 'Como funciona então?',
                description:
                    ' O Pacstall recebe ficheiros chamados de <0>pacscripts</0> (semelhantes ' +
                    ' ao PKGBUILD) que contêm a informação necessária para compilar pacotes ' +
                    ' e compila-os como executáveis no seu sistema.',
            },
        },
        installationInstructions: 'Instruções de Instalação',
        showcase: {
            title: 'Demonstração',
            packageSearch: 'Procurar Pacotes',
        },
    },
    navbar: {
        title: 'Pacstall',
        contribute: {
            title: 'Contribua',
            workOnFeatures: 'Trabalhe em novas funcionalidades',
            helpTranslate: 'Ajude com as traduções',
            becomeAMaintainer: 'Torne-se um responsável de pacotes',
        },
        social: {
            title: 'Redes Sociais',
            discord: 'Discord',
            matrix: 'Matrix',
            reddit: 'Reddit',
            lemmy: 'Lemmy',
            mastodon: 'Mastodon',
        },
        browse: {
            title: 'Navegar pelos pacotes',
        },
        privacy: {
            title: 'Política de Privacidade',
        },
        install: 'Instalar',
    },
    cookieConsent: {
        title: 'Aviso sobre Cookies (TL;DR)',
        paragraphs: [
            'Olá, estamos a utilizar cookies na nosso website.',
            'Não gostamos de lhe dar informações enganadoras ou confusas. Apenas utilizamos cookies para funcionalidades essenciais, como definições de temas, localização e autenticação.',
            'Pode ler a política de privacidade completa <0>aqui <1/></0>.',
            'Ao continuar a utilizar este site, está a <strong>concordar com a política de privacidade</strong>.',
        ],
        accept: 'Aceitar',
    },
    packageSearch: {
        dropdown: {
            package: 'Pacote',
            packageTooltip: 'Procurar em nomes e descrições de pacotes',
            maintainer: 'Responsável',
            maintainerTooltip: 'Procurar por nomes e e-mails de responsáveis',
        },
        table: {
            name: 'Nome',
            maintainer: 'Responsável',
            version: 'Versão',
            install: 'Instalar',
        },
        versionTooltip: {
            notInRegistry: 'Este pacote não está registado no Repology',
            latest: 'Esta é a versão mais recente do pacote',
            patch: 'Este pacote tem uma atualização de correção disponível',
            minor: 'Este pacote tem uma pequena atualização disponível',
            major: 'Este pacote tem uma grande atualização disponível',
            isGit: 'Este pacote é um pacote Git',
        },
        noResults: 'Não encontrou o que procurava? <0>Crie um pedido!</0>',
        search: 'Procurar',
        orphaned: 'Órfão',
        maintainerTooltip: {
            maintainedBy: 'Este pacote está a ser mantido por {{ name }}',
            noMaintainer: 'Este pacote não está a ser mantido',
        },
        pagination: {
            previous: 'anterior',
            next: 'seguinte',
        },
    },
    packageDetails: {
        table: {
            name: 'Nome',
            version: 'Versão',
            maintainer: 'Responsável',
            dependencies: 'Dependências',
            requiredBy: 'Requerido Por',
        },
        orphaned: 'Órfão',
        noResults: 'Nenhum',
        openInGithub: 'Abrir no Github',
        view: 'Visualizar',
        howToInstall: {
            title: 'Como Instalar',
            step1: '1º Passo: Configurar o Pacstall',
            step2: '2º Passo: Instalar {{ name }}',
        },
        dependenciesModal: {
            title: 'Dependências',
            buildDependencies: 'Dependências de compilação',
            optionalDependencies: 'Dependências opcionais',
            runtimeDependencies: 'Dependências em tempo de execução',
            pacstallDependencies: 'Dependências de Pacstall',
            name: 'Nome',
            close: 'Fechar',
            provider: 'Provedor',
            noDescription: 'Nenhuma descrição disponível',
            version: 'Versão',
        },
        requiredByModal: {
            title: 'Requerido por',
            name: 'Nome',
            provider: 'Provedor',
            close: 'Fechar',
            noDescription: 'Nenhuma descrição disponível',
        },
    },
}
