import Locale from './locale'

export default <Locale>{
    home: {
        title: 'প্যাকস্টল',
        subtitle: 'উবুন্টুর জন্য আদর্শ এইয়ুআর (AUR)',
        cards: {
            whyDifferent: {
                title: 'এটি কিভাবে অন্যান্য প্যাকেজ ম্যানেজারের থেকে আলাদা?',
                description:
                    'প্যাকস্টল আপনাকে উবুন্টুর স্টেবেল বেসের সাথে আপোস না করেই ব্লিডিং এজ সফ্টওয়্যার ব্যবহার করার' +
                    'সুযোগ দেয়, তাই আপনাকে কোনো সফটওয়্যারের সুরক্ষা প্যাচ বা নতুন ফিচারের বিষয়ে চিন্তা করার প্রয়োজন নেই।',
            },
            howItWorks: {
                title: 'তাহলে এটা কিভাবে কাজ করে?',
                description:
                    'প্যাকস্টল <0>প্যাকস্ক্রিপ্টস</0> (প্যাকেজ বিল্ড এর মত) ' +
                    'নামের ফাইল ব্যবহার করে যাতে প্যাকেজ তৈরির জন্য প্রয়োজনীয় বিষয়বস্তু থাকে এবং সেগুলো থেকে ' +
                    'আপনার সিস্টেমে একটি একজিকিউটেবল তৈরি করে।',
            },
        },
        installationInstructions: 'ইন্সটলেশনের নির্দেশ',
        showcase: {
            title: 'প্রদর্শনী',
            packageSearch: 'প্যাকেজ অনুসন্ধান',
        },
    },
    navbar: {
        title: 'প্যাকস্টল',
        contribute: {
            title: 'প্রকল্পে অবদান',
            workOnFeatures: 'নতুন ফিচার তৈরি করতে সাহায্য করুন',
            helpTranslate: 'এই ওয়েবসাইটটিকে অনুবাদ করতে সাহায্য করুন',
            becomeAMaintainer: 'একটি প্যাকস্ক্রিপ্ট সম্প্রণেতা হয়ে উঠুন',
        },
        social: {
            title: 'সামাজিক মাধ্যম',
            discord: 'ডিসকর্ড',
            matrix: 'ম্যাট্রিক্স',
            reddit: 'রেড্ডিট',
            lemmy: 'লেমী',
            mastodon: 'ম্যাস্টোডন',
        },
        browse: {
            title: 'প্যাকেজগুলো ব্রাউজ করুন',
        },
        privacy: {
            title: 'গোপনীয়তা নীতি',
        },
        install: 'ইনস্টল',
    },
    cookieConsent: {
        title: 'কুকির বিজ্ঞপ্তি',
        paragraphs: [
            'নমস্কার, হ্যাঁ আমরা কুকি ব্যবহার করি।',
            'আমরা আপনাকে বিভ্রান্তিকর তথ্য দিতে চাই না। আমরা শুধুমাত্র প্রয়োজনীয় ফিচারস যেমন থিম সেটিংস, স্থানীয়করণ এবং প্রমাণীকরণের জন্য কুকিস ব্যবহার করি।',
            'আপনি আমাদের সম্পূর্ণ গোপনীয়তা নীতি <0>এখানে<1/></0> পড়তে পারেন।',
            'এই সাইটটি ব্যবহার চালিয়ে যাওয়ার মাধ্যমে, আপনি আমদের <strong>গোপনীয়তা নীতিতে সম্মতি দিচ্ছেন</strong>।',
        ],
        accept: 'ঠিক আছে',
    },
    packageSearch: {
        dropdown: {
            package: 'প্যাকেজ',
            packageTooltip: 'প্যাকেজ গুলোর নামে এবং বিবরণে অনুসন্ধান করে',
            maintainer: 'সম্প্রণেতা',
            maintainerTooltip:
                'সম্প্রণেতার নামের অথবা ইমেলের দ্বারা অনুসন্ধান করে',
        },
        table: {
            name: 'নাম',
            maintainer: 'সম্প্রণেতা',
            version: 'সংস্করণ',
            install: 'ইনস্টল',
        },
        versionTooltip: {
            notInRegistry: 'এই প্যাকেজটি রিপোলজি রেজিস্ট্রিতে নেই',
            latest: 'এই প্যাকেজটি অধুনা সংস্করণে রয়েছে',
            hasPatchUpdate: 'এই প্যাকেজটির একটি  প্যাচ আপডেট অপেক্ষারত রয়েছে',
            hasMinorUpdate: 'এই প্যাকেজটির একটি  লঘু আপডেট অপেক্ষারত রয়েছে',
            hasMajorUpdate: 'এই প্যাকেজটির একটি  গুরু আপডেট অপেক্ষারত রয়েছে',
            isGit: 'এই প্যাকেজটি একটি গিট প্যাকেজ',
        },
        noResults: 'আপনি যা খুঁজছেন তা পাচ্ছেন না? <0>অনুরোধ রাখুন!</0>',
        search: 'অনুসন্ধান করুন',
        orphaned: 'সম্প্রণেতা হীন',
        maintainerTooltip: {
            maintainedBy: 'এই প্যাকেজটি {{ name }} সম্প্রণ করছেন',
            noMaintainer: 'এই প্যাকেজটি সম্প্রণ করা হচ্ছে না',
        },
        pagination: {
            previous: 'আগে',
            next: 'পরে',
        },
    },
    packageDetails: {
        table: {
            name: 'নাম',
            version: 'সংস্করণ',
            maintainer: 'সম্প্রণেতা',
            dependencies: 'নির্ভরতা',
            requiredBy: 'অন্বিষ্ট',
            lastUpdatedAt: 'শেষ আপডেট',
        },
        orphaned: 'সম্প্রণেতা হীন',
        noResults: 'শূন্য',
        openInGithub: 'গিটহাবে খুলুন',
        view: 'দেখুন',
        howToInstall: {
            title: 'কিভাবে ইনস্টল করতে হবে',
            step1: 'ধাপ ১: প্যাকস্টল ইনস্টল করুন',
            step2: 'ধাপ ২: {{ name }} ইনস্টল করুন',
        },
        dependenciesModal: {
            title: 'নির্ভরতা',
            buildDependencies: 'গঠক নির্ভরতা',
            optionalDependencies: 'ঐচ্ছিক নির্ভরতা',
            runtimeDependencies: 'চলক নির্ভরতা',
            pacstallDependencies: 'প্যাকস্টল নির্ভরতা',
            name: 'নাম',
            close: 'বন্ধ',
            provider: 'প্রদানকারী',
            noDescription: 'কোনো বর্ণনা নেই',
            version: 'সংস্করণ',
        },
        requiredByModal: {
            title: 'অন্বিষ্ট',
            name: 'নাম',
            provider: 'প্রদানকারী',
            close: 'বন্ধ',
            noDescription: 'কোনো বর্ণনা নেই',
        },
    },
}
