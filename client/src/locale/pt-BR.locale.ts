import Locale from './locale'

export default <Locale>{
    home: {
        title: 'Pacstall',
        subtitle: 'O AUR para Ubuntu',
        cards: {
            whyDifferent: {
                title: 'Por que isso é diferente de qualquer outro gerenciador de pacotes?',
                description:
                    'O Pacstall utiliza a base estável do Ubuntu, mas permite que você use ' +
                    'software de última geração com poucos ou nenhum compromisso, sem que ' +
                    'você precise se preocupar com correções de segurança ou novos recursos.',
            },
            howItWorks: {
                title: 'How does it work then?',
                description:
                    ' O Pacstall recebe arquivos chamados de <0>pacscripts</0> (semelhantes ' +
                    ' ao PKGBUILD) que contêm a informação necessária para compilar pacotes ' +
                    ' e os compila como executáveis no seu sistema.',
            },
        },
        installationInstructions: 'Instruções de Instalação',
        showcase: {
            title: 'Demonstração',
            packageSearch: 'Buscar Pacotes',
        },
    },
    navbar: {
        title: 'Pacstall',
        contribute: {
            title: 'Contribua',
            workOnFeatures: 'Trabalhe em novas funcionalidades',
            helpTranslate: 'Ajude com traduções',
            becomeAMaintainer: 'Torne-se um mantenedor de pacotes',
        },
        social: {
            title: 'Redes Sociais',
            discord: 'Discord',
            matrix: 'Matrix',
            reddit: 'Reddit',
            mastodon: 'Mastodon',
        },
        browse: {
            title: 'Navegar pelos Pacotes',
        },
        privacy: {
            title: 'Poliítica de Privacidade',
        },
        install: 'Instalar',
    },
    cookieConsent: {
        title: 'Aviso sobre Cookies (TL;DR)',
        paragraphs: [
            'Olá, sim, nós utilizamos cookies.',
            'Não gostamos de fornecer informações enganosas ou confusas. Nós utilizamos cookies apenas para recursos essenciais, como configurações de tema, localização e autenticação.',
            'Você pode ler a política de privacidade completa <0>aqui <1/></0>.',
            'Ao continuar a usar este site, você <strong>concorda com a política de privacidade</strong>.',
        ],
        accept: 'ok, legal',
    },
    packageSearch: {
        dropdown: {
            package: 'Pacote',
            packageTooltip: 'Busca em nomes e descrições de pacotes',
            maintainer: 'Mantenedor',
            maintainerTooltip: 'Busca por nomes e e-mails de mantenedores',
        },
        table: {
            name: 'Nome',
            maintainer: 'Mantenedor',
            version: 'Versão',
            install: 'Instalar',
        },
        versionTooltip: {
            notInRegistry: 'Este pacote não está registrado no Repology',
            latest: 'Esta é a versão mais recente do pacote',
            patch:
                'Este pacote possui uma atualização de correção disponível',
            minor:
                'Este pacote possui uma atualização menor disponível',
            major:
                'Este pacote possui uma atualização maior disponível',
            isGit: 'Este pacote é um pacote Git',
        },
        noResults: 'Não encontrou o que buscava? <0>Faça uma solicitação!</0>',
        search: 'Buscar',
        orphaned: 'Órfão',
        maintainerTooltip: {
            maintainedBy: 'Este pacote está sendo mantido por {{ name }}',
            noMaintainer: 'Este pacote não está sendo mantido',
        },
        pagination: {
            previous: 'anterior',
            next: 'próximo',
        },
    },
    packageDetails: {
        table: {
            name: 'Nome',
            version: 'Versão',
            maintainer: 'Mantenedor',
            dependencies: 'Dependências',
            requiredBy: 'Requerido Por',
        },
        orphaned: 'Órfão',
        noResults: 'Nenhum',
        openInGithub: 'Abrir No Github',
        view: 'Visualizar',
        howToInstall: {
            title: 'Como Instalar',
            step1: '1º Passo: Configurar o Pacstall',
            step2: '2º Passo: Instalar {{ name }}',
        },
        dependenciesModal: {
            title: 'Dependências',
            buildDependencies: 'Dependências de Compilação',
            optionalDependencies: 'Dependências Opcionais',
            runtimeDependencies: 'Dependências em Tempo de Execução',
            pacstallDependencies: 'Dependências de Pacstall',
            name: 'Nome',
            close: 'Fechar',
            provider: 'Provedor',
            noDescription: 'Nenhuma descrição disponível',
        },
        requiredByModal: {
            title: 'Requerido Por',
            name: 'Nome',
            provider: 'Provedor',
            close: 'Fechar',
            noDescription: 'Nenhuma descrição disponível',
        },
    },
}
