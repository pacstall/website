import Locale from './locale'

export default <Locale>{
    home: {
        title: 'Pacstall',
        subtitle: 'Ubuntu 上的 AUR',
        cards: {
            whyDifferent: {
                title: '它与其他包管理器有何不同？',
                description:
                    'Pacstall 以稳定的 Ubuntu 为基础，同时让您能毫不妥协地' +
                    '使用最新软件，因此您无需再为安全补丁或新功能而操心。',
            },
            howItWorks: {
                title: '它是如何工作的？',
                description:
                    ' Pacstall 会读取一种名为 <0>pacscripts</0> 的文件（类似于' +
                    ' PKGBUILD），它包含了构建软件包所需的信息。随后，Pacstall' +
                    ' 会在您的系统上将这些软件包构建为可执行文件。',
            },
        },
        installationInstructions: '安装指南',
        showcase: {
            title: '使用展示',
            packageSearch: '搜索软件包',
        },
    },
    navbar: {
        title: 'Pacstall',
        contribute: {
            title: '贡献',
            workOnFeatures: '参与功能开发',
            helpTranslate: '协助翻译工作',
            becomeAMaintainer: '成为软件包维护者',
        },
        social: {
            title: '社交网络',
            discord: 'Discord',
            matrix: 'Matrix',
            reddit: 'Reddit',
            lemmy: 'Lemmy',
            mastodon: 'Mastodon',
        },
        browse: {
            title: '浏览软件包',
        },
        privacy: {
            title: '隐私政策',
        },
        install: '安装',
    },
    cookieConsent: {
        title: 'Cookie 声明简述',
        paragraphs: [
            '您好，本网站会使用 Cookie。',
            '我们不会提供任何误导性或令人困惑的信息。我们仅将 Cookie 用于主题切换、语言本地化和用户认证等必要功能。',
            '您可以<0>点此</0>阅读完整的隐私政策<1/>。',
            '若您继续使用本网站，即表示您<strong>同意本隐私政策</strong>。',
        ],
        accept: '我明白了',
    },
    packageSearch: {
        dropdown: {
            package: '软件包',
            packageTooltip: '按软件包名称或描述搜索',
            maintainer: '维护者',
            maintainerTooltip: '按维护者名称或邮箱搜索',
        },
        table: {
            name: '名称',
            maintainer: '维护者',
            version: '版本',
            install: '安装',
        },
        versionTooltip: {
            notInRegistry: '此软件包未收录于 Repology',
            latest: '已是最新版本',
            patch: '有可用的补丁更新',
            minor: '有可用的次要版本更新',
            major: '有可用的主要版本更新',
            isGit: '这是一个 Git 软件包',
        },
        noResults: '没找到您想要的？ <0>提交一个打包请求！</0>',
        search: '搜索',
        orphaned: '无人维护',
        maintainerTooltip: {
            maintainedBy: '此包由 {{ name }} 维护',
            noMaintainer: '此包当前无人维护',
        },
        pagination: {
            previous: '上一页',
            next: '下一页',
        },
    },
    packageDetails: {
        table: {
            name: '名称',
            version: '版本',
            maintainer: '维护者',
            dependencies: '依赖',
            requiredBy: '被依赖',
            lastUpdatedAt: '最后更新',
        },
        orphaned: '无人维护',
        noResults: '无',
        openInGithub: '在 GitHub 上查看',
        view: '查看',
        howToInstall: {
            title: '安装步骤',
            step1: '第 1 步：配置 Pacstall',
            step2: '第 2 步：安装 {{ name }}',
        },
        dependenciesModal: {
            title: '依赖',
            buildDependencies: '构建依赖',
            optionalDependencies: '可选依赖',
            runtimeDependencies: '运行时依赖',
            pacstallDependencies: 'Pacstall 依赖',
            name: '名称',
            close: '关闭',
            provider: '提供者',
            noDescription: '暂无描述',
            version: '版本',
        },
        requiredByModal: {
            title: '被依赖',
            name: '名称',
            provider: '提供者',
            close: '关闭',
            noDescription: '暂无描述',
        },
    },
}
