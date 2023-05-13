import Locale from './locale'

export default <Locale>{
    home: {
        title: 'প্যাকস্টল',
        subtitle: 'উবুন্টুর জন্য আদর্শ এইউআর (AUR)',
        cards: {
            whyDifferent: {
                title: 'এটি কিভাবে অন্যান্য প্যাকেজ ম্যানেজারের থেকে আলাদা?',
                description:
                'প্যাকস্টল আপনাকে উবুন্টুর স্থিতিশীলতা সাথে আপস না করেই ব্লিডিং এজ সফ্টওয়্যার ব্যবহার করার' +
                    "ক্ষমতা দেয়, তাই আপনাকে কন সফটওয়্যারের সুরক্ষা প্যাচ বা নতুন ফিচারের বিষয়ে চিন্তা করতে হবে না।",
            },
            howItWorks: {
                title: 'তাহলে এটা কিভাবে কাজ করে?',
                description:
    		    'প্যাকস্টল <0>প্যাকস্ক্রিপ্টস (pacscripts)</0> (পিকেজিবিল্ডস (PKGBUILDs) এর মত) ' +
    		        "নামের ফাইল ব্যবহার করে যাতে প্যাকেজ তৈরির জন্য প্রয়োজনীয় বিষয়বস্তু থাকে এবং সেগুলো থেকে " +
    		        'আপনার সিস্টেমে একটি একজিকিউটেবল তৈরি করে।'
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
            title: 'আপনার অবদান দিন',
            workOnFeatures: 'নতুন ফিচার তৈরি করতে সাহায্য করুন',
            helpTranslate: 'এই ওয়েবসাইটটিকে তর্জমা করতে সাহায্য করুন',
            becomeAMaintainer: 'একটি প্যাকেজ রক্ষণাবেক্ষণকারী হয়ে উঠুন',
        },
        social: {
            title: 'সামাজিক যোগাযোগ',
            discord: 'ডিসকর্ড',
            matrix: 'ম্যাট্রিক্স',
            reddit: 'রেড্ডিট',
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
            'নমস্কার, হ্যাঁ আমরা কুকিজ ব্যবহার করি।',
            "আমরা আপনাকে বিভ্রান্তিকর তথ্য দিতে পছন্দ করি না। আমরা শুধুমাত্র প্রয়োজনীয় বৈশিষ্ট্য যেমন থিম সেটিংস, স্থানীয়করণ এবং প্রমাণীকরণের জন্য কুকি ব্যবহার করি।",
            'আপনি আমাদের সম্পূর্ণ গোপনীয়তা নীতি <0>এখানে <1/></0> পড়তে পারেন।',
            'এই সাইটটি ব্যবহার করা চালিয়ে যাওয়ার মাধ্যমে, আপনি <strong>গোপনীয়তা নীতিতে আপনার সম্মতি দেন</strong>।',
        ],
        accept: 'ঠিক আছে',
    },
    packageSearch: {
        dropdown: {
            package: 'প্যাকেজ',
            packageTooltip: 'প্যাকেজ গুলোর নামে এবং বিবরণে অনুসন্ধান করে',
            maintainer: 'রক্ষণাবেক্ষণকারী',
            maintainerTooltip: 'রক্ষণাবেক্ষণকারীর নামের অথবা ইমেলের দ্বারা অনুসন্ধান করে',
        },
        table: {
            name: 'নাম',
            maintainer: 'রক্ষণাবেক্ষণকারী',
            version: 'সংস্করণ',
            install: 'ইনস্টল',
        },
        versionTooltip: {
            notInRegistry: 'এই প্যাকেজটি রিপোলজি রেজিস্ট্রিতে নেই',
            latest: 'এই প্যাকেজটি সাম্প্রতিকতম সংস্করণে রয়েছে',
            hasPatchUpdate: 'এই প্যাকেজটির একটি  প্যাচ আপডেট অপেক্ষারত রয়েছ',
            hasMinorUpdate: 'এই প্যাকেজটির একটি  মিনর আপডেট অপেক্ষারত রয়েছ',
            hasMajorUpdate: 'এই প্যাকেজটির একটি  মেজর আপডেট অপেক্ষারত রয়েছ',
            isGit: 'এই প্যাকেজটি একটি গিট প্যাকেজ',
        },
        noResults: 'যা খুঁজছেন তা কি পাচ্ছেন না? <0>অনুরোধ করুন!</0>',
        search: 'অনুসন্ধান করুন',
        orphaned: 'রক্ষণাবেক্ষণকারী হীন',
        maintainerTooltip: {
            maintainedBy: 'এই প্যাকেজটি {{ name }} রক্ষণাবেক্ষণ করা হচ্ছে',
            noMaintainer: 'এই প্যাকেজটি রক্ষণাবেক্ষণ করা হচ্ছে না',
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
            maintainer: 'রক্ষণাবেক্ষণকারী',
            dependencies: 'নির্ভরতা',
            requiredBy: 'দ্বারা প্রয়োজনীয়', 
        },
        orphaned: 'রক্ষণাবেক্ষণকারী হীন',
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
            buildDependencies: 'গঠন নির্ভরশীলতা',
            optionalDependencies: 'ঐচ্ছিক নির্ভরশীলতা',
            runtimeDependencies: 'চলাকালীন নির্ভরশীলতা',
            pacstallDependencies: 'প্যাকস্টল নির্ভরশীলতা',
            name: 'নাম',
            close: 'বন্ধ',
            provider: 'প্রদানকারী',
            noDescription: 'কোন বর্ণনা নেই',
        },
        requiredByModal: {
            title: 'দ্বারা প্রয়োজনীয়', // Likely to be incorrect, word order needs to change for accurate translation.
            name: 'নাম',
            provider: 'প্রদানকারী',
            close: 'বন্ধ',
            noDescription: 'কোন বর্ণনা নেই',
        },
    },
}
