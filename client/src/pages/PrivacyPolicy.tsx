import {
    Container,
    Heading,
    HStack,
    Link,
    ListItem,
    Stack,
    StackProps,
    Text,
    UnorderedList,
    useBreakpointValue,
} from '@chakra-ui/react'
import { FC } from 'react'
import { Helmet } from 'react-helmet'
import PageAnimation from '../components/animations/PageAnimation'

const ResponsiveStack: FC<StackProps> = props => {
    const UsedStack = useBreakpointValue({ base: Stack, md: HStack })
    return <UsedStack {...props} />
}

const PrivacyPolicy: FC = () => (
    <>
        <Helmet>
            <title>Privacy Policy - Pacstall</title>
        </Helmet>

        <PageAnimation>
            <Container my='10' maxW='60em' textAlign='justify'>
                <ResponsiveStack mb='5' justifyContent='space-between'>
                    <Heading textAlign='left'>Privacy Policy</Heading>
                    <Text>Last updated: March 04, 2022</Text>
                </ResponsiveStack>

                <Text>
                    This Privacy Policy describes Our policies and procedures on
                    the collection, use and disclosure of Your information when
                    You use the Service and tells You about Your privacy rights
                    and how the law protects You.
                </Text>
                <Text>
                    We use Your Personal data to provide and improve the
                    Service. By using the Service, You agree to the collection
                    and use of information in accordance with this Privacy
                    Policy. This Privacy Policy has been created with the help
                    of the{' '}
                    <a
                        href='https://www.privacypolicies.com/blog/privacy-policy-template/'
                        target='_blank'
                    >
                        Privacy Policy Template
                    </a>
                    .
                </Text>
                <Heading textAlign='left' mb='3' mt='7'>
                    Interpretation and Definitions
                </Heading>
                <Heading textAlign='left' size='md' mt='3' mb='2'>
                    Interpretation
                </Heading>
                <Text>
                    The words of which the initial letter is capitalized have
                    meanings defined under the following conditions. The
                    following definitions shall have the same meaning regardless
                    of whether they appear in singular or in plural.
                </Text>
                <Heading textAlign='left' size='md' mt='3' mb='2'>
                    Definitions
                </Heading>
                <Text>For the purposes of this Privacy Policy:</Text>
                <UnorderedList>
                    <ListItem>
                        <Text>
                            <strong>Account</strong> means a unique account
                            created for You to access our Service or parts of
                            our Service.
                        </Text>
                    </ListItem>
                    <ListItem>
                        <Text>
                            <strong>Affiliate</strong> means an entity that
                            controls, is controlled by or is under common
                            control with a party, where &quot;control&quot;
                            means ownership of 50% or more of the shares, equity
                            interest or other securities entitled to vote for
                            election of directors or other managing authority.
                        </Text>
                    </ListItem>
                    <ListItem>
                        <Text>
                            <strong>Application</strong> means the software
                            program provided by the Company downloaded by You on
                            any electronic device, named Pacstall
                        </Text>
                    </ListItem>
                    <ListItem>
                        <Text>
                            <strong>Company</strong> (referred to as either
                            &quot;the Company&quot;, &quot;We&quot;,
                            &quot;Us&quot; or &quot;Our&quot; in this Agreement)
                            refers to Pacstall.
                        </Text>
                    </ListItem>
                    <ListItem>
                        <Text>
                            <strong>Country</strong> refers to: Washington,
                            United States
                        </Text>
                    </ListItem>
                    <ListItem>
                        <Text>
                            <strong>Device</strong> means any device that can
                            access the Service such as a computer, a cellphone
                            or a digital tablet.
                        </Text>
                    </ListItem>
                    <ListItem>
                        <Text>
                            <strong>Personal Data</strong> is any information
                            that relates to an identified or identifiable
                            individual.
                        </Text>
                    </ListItem>
                    <ListItem>
                        <Text>
                            <strong>Service</strong> refers to the Application.
                        </Text>
                    </ListItem>
                    <ListItem>
                        <Text>
                            <strong>Service Provider</strong> means any natural
                            or legal person who processes the data on behalf of
                            the Company. It refers to third-party companies or
                            individuals employed by the Company to facilitate
                            the Service, to provide the Service on behalf of the
                            Company, to perform services related to the Service
                            or to assist the Company in analyzing how the
                            Service is used.
                        </Text>
                    </ListItem>
                    <ListItem>
                        <Text>
                            <strong>Usage Data</strong> refers to data collected
                            automatically, either generated by the use of the
                            Service or from the Service infrastructure itself
                            (for example, the duration of a page visit).
                        </Text>
                    </ListItem>
                    <ListItem>
                        <Text>
                            <strong>You</strong> means the individual accessing
                            or using the Service, or the company, or other legal
                            entity on behalf of which such individual is
                            accessing or using the Service, as applicable.
                        </Text>
                    </ListItem>
                </UnorderedList>
                <Heading textAlign='left' mb='3' mt='7'>
                    Collecting and Using Your Personal Data
                </Heading>
                <Heading textAlign='left' size='md' mt='3' mb='2'>
                    Types of Data Collected
                </Heading>
                <Heading textAlign='left' size='sm' mt='2' mb='1'>
                    Personal Data
                </Heading>
                <Text>
                    While using Our Service, We may ask You to provide Us with
                    certain personally identifiable information that can be used
                    to contact or identify You. Personally identifiable
                    information may include, but is not limited to:
                </Text>
                <UnorderedList>
                    <ListItem>Usage Data</ListItem>
                </UnorderedList>
                <Heading textAlign='left' size='sm' mt='2' mb='1'>
                    Usage Data
                </Heading>
                <Text>
                    Usage Data is collected automatically when using the
                    Service.
                </Text>
                <Text>
                    Usage Data may include information such as Your Device's
                    Internet Protocol address (e.g. IP address), browser type,
                    browser version, the pages of our Service that You visit,
                    the time and date of Your visit, the time spent on those
                    pages, unique device identifiers and other diagnostic data.
                </Text>
                <Text>
                    When You access the Service by or through a mobile device,
                    We may collect certain information automatically, including,
                    but not limited to, the type of mobile device You use, Your
                    mobile device unique ID, the IP address of Your mobile
                    device, Your mobile operating system, the type of mobile
                    Internet browser You use, unique device identifiers and
                    other diagnostic data.
                </Text>
                <Text>
                    We may also collect information that Your browser sends
                    whenever You visit our Service or when You access the
                    Service by or through a mobile device.
                </Text>
                <Heading textAlign='left' size='md' mt='3' mb='2'>
                    Use of Your Personal Data
                </Heading>
                <Text>
                    The Company may use Personal Data for the following
                    purposes:
                </Text>
                <UnorderedList>
                    <ListItem>
                        <Text>
                            <strong>To provide and maintain our Service</strong>
                            , including to monitor the usage of our Service.
                        </Text>
                    </ListItem>
                    <ListItem>
                        <Text>
                            <strong>To manage Your Account:</strong> to manage
                            Your registration as a user of the Service. The
                            Personal Data You provide can give You access to
                            different functionalities of the Service that are
                            available to You as a registered user.
                        </Text>
                    </ListItem>
                    <ListItem>
                        <Text>
                            <strong>For the performance of a contract:</strong>{' '}
                            the development, compliance and undertaking of the
                            purchase contract for the products, items or
                            services You have purchased or of any other contract
                            with Us through the Service.
                        </Text>
                    </ListItem>
                    <ListItem>
                        <Text>
                            <strong>To contact You:</strong> To contact You by
                            email, telephone calls, SMS, or other equivalent
                            forms of electronic communication, such as a mobile
                            application's push notifications regarding updates
                            or informative communications related to the
                            functionalities, products or contracted services,
                            including the security updates, when necessary or
                            reasonable for their implementation.
                        </Text>
                    </ListItem>
                    <ListItem>
                        <Text>
                            <strong>To provide You</strong> with news, special
                            offers and general information about other goods,
                            services and events which we offer that are similar
                            to those that you have already purchased or enquired
                            about unless You have opted not to receive such
                            information.
                        </Text>
                    </ListItem>
                    <ListItem>
                        <Text>
                            <strong>To manage Your requests:</strong> To attend
                            and manage Your requests to Us.
                        </Text>
                    </ListItem>
                    <ListItem>
                        <Text>
                            <strong>For business transfers:</strong> We may use
                            Your information to evaluate or conduct a merger,
                            divestiture, restructuring, reorganization,
                            dissolution, or other sale or transfer of some or
                            all of Our assets, whether as a going concern or as
                            part of bankruptcy, liquidation, or similar
                            proceeding, in which Personal Data held by Us about
                            our Service users is among the assets transferred.
                        </Text>
                    </ListItem>
                    <ListItem>
                        <Text>
                            <strong>For other purposes</strong>: We may use Your
                            information for other purposes, such as data
                            analysis, identifying usage trends, determining the
                            effectiveness of our promotional campaigns and to
                            evaluate and improve our Service, products,
                            services, marketing and your experience.
                        </Text>
                    </ListItem>
                </UnorderedList>
                <Text>
                    We may share Your personal information in the following
                    situations:
                </Text>
                <UnorderedList>
                    <ListItem>
                        <strong>With Service Providers:</strong> We may share
                        Your personal information with Service Providers to
                        monitor and analyze the use of our Service, to contact
                        You.
                    </ListItem>
                    <ListItem>
                        <strong>For business transfers:</strong> We may share or
                        transfer Your personal information in connection with,
                        or during negotiations of, any merger, sale of Company
                        assets, financing, or acquisition of all or a portion of
                        Our business to another company.
                    </ListItem>
                    <ListItem>
                        <strong>With Affiliates:</strong> We may share Your
                        information with Our affiliates, in which case we will
                        require those affiliates to honor this Privacy Policy.
                        Affiliates include Our parent company and any other
                        subsidiaries, joint venture partners or other companies
                        that We control or that are under common control with
                        Us.
                    </ListItem>
                    <ListItem>
                        <strong>With business partners:</strong> We may share
                        Your information with Our business partners to offer You
                        certain products, services or promotions.
                    </ListItem>
                    <ListItem>
                        <strong>With other users:</strong> when You share
                        personal information or otherwise interact in the public
                        areas with other users, such information may be viewed
                        by all users and may be publicly distributed outside.
                    </ListItem>
                    <ListItem>
                        <strong>With Your consent</strong>: We may disclose Your
                        personal information for any other purpose with Your
                        consent.
                    </ListItem>
                </UnorderedList>
                <Heading textAlign='left' size='md' mt='3' mb='2'>
                    Retention of Your Personal Data
                </Heading>
                <Text>
                    The Company will retain Your Personal Data only for as long
                    as is necessary for the purposes set out in this Privacy
                    Policy. We will retain and use Your Personal Data to the
                    extent necessary to comply with our legal obligations (for
                    example, if we are required to retain your data to comply
                    with applicable laws), resolve disputes, and enforce our
                    legal agreements and policies.
                </Text>
                <Text>
                    The Company will also retain Usage Data for internal
                    analysis purposes. Usage Data is generally retained for a
                    shorter period of time, except when this data is used to
                    strengthen the security or to improve the functionality of
                    Our Service, or We are legally obligated to retain this data
                    for longer time periods.
                </Text>
                <Heading textAlign='left' size='md' mt='3' mb='2'>
                    Transfer of Your Personal Data
                </Heading>
                <Text>
                    Your information, including Personal Data, is processed at
                    the Company's operating offices and in any other places
                    where the parties involved in the processing are located. It
                    means that this information may be transferred to — and
                    maintained on — computers located outside of Your state,
                    province, country or other governmental jurisdiction where
                    the data protection laws may differ than those from Your
                    jurisdiction.
                </Text>
                <Text>
                    Your consent to this Privacy Policy followed by Your
                    submission of such information represents Your agreement to
                    that transfer.
                </Text>
                <Text>
                    The Company will take all steps reasonably necessary to
                    ensure that Your data is treated securely and in accordance
                    with this Privacy Policy and no transfer of Your Personal
                    Data will take place to an organization or a country unless
                    there are adequate controls in place including the security
                    of Your data and other personal information.
                </Text>
                <Heading textAlign='left' size='md' mt='3' mb='2'>
                    Disclosure of Your Personal Data
                </Heading>
                <Heading textAlign='left' size='sm' mt='2' mb='1'>
                    Business Transactions
                </Heading>
                <Text>
                    If the Company is involved in a merger, acquisition or asset
                    sale, Your Personal Data may be transferred. We will provide
                    notice before Your Personal Data is transferred and becomes
                    subject to a different Privacy Policy.
                </Text>
                <Heading textAlign='left' size='sm' mt='2' mb='1'>
                    Law enforcement
                </Heading>
                <Text>
                    Under certain circumstances, the Company may be required to
                    disclose Your Personal Data if required to do so by law or
                    in response to valid requests by public authorities (e.g. a
                    court or a government agency).
                </Text>
                <Heading textAlign='left' size='sm' mt='2' mb='1'>
                    Other legal requirements
                </Heading>
                <Text>
                    The Company may disclose Your Personal Data in the good
                    faith belief that such action is necessary to:
                </Text>
                <UnorderedList>
                    <ListItem>Comply with a legal obligation</ListItem>
                    <ListItem>
                        Protect and defend the rights or property of the Company
                    </ListItem>
                    <ListItem>
                        Prevent or investigate possible wrongdoing in connection
                        with the Service
                    </ListItem>
                    <ListItem>
                        Protect the personal safety of Users of the Service or
                        the public
                    </ListItem>
                    <ListItem>Protect against legal liability</ListItem>
                </UnorderedList>
                <Heading textAlign='left' size='md' mt='3' mb='2'>
                    Security of Your Personal Data
                </Heading>
                <Text>
                    The security of Your Personal Data is important to Us, but
                    remember that no method of transmission over the Internet,
                    or method of electronic storage is 100% secure. While We
                    strive to use commercially acceptable means to protect Your
                    Personal Data, We cannot guarantee its absolute security.
                </Text>
                <Heading textAlign='left' mb='3' mt='7'>
                    Children's Privacy
                </Heading>
                <Text>
                    Our Service does not address anyone under the age of 13. We
                    do not knowingly collect personally identifiable information
                    from anyone under the age of 13. If You are a parent or
                    guardian and You are aware that Your child has provided Us
                    with Personal Data, please contact Us. If We become aware
                    that We have collected Personal Data from anyone under the
                    age of 13 without verification of parental consent, We take
                    steps to remove that information from Our servers.
                </Text>
                <Text>
                    If We need to rely on consent as a legal basis for
                    processing Your information and Your country requires
                    consent from a parent, We may require Your parent's consent
                    before We collect and use that information.
                </Text>
                <Heading textAlign='left' mb='3' mt='7'>
                    Links to Other Websites
                </Heading>
                <Text>
                    Our Service may contain links to other websites that are not
                    operated by Us. If You click on a third party link, You will
                    be directed to that third party's site. We strongly advise
                    You to review the Privacy Policy of every site You visit.
                </Text>
                <Text>
                    We have no control over and assume no responsibility for the
                    content, privacy policies or practices of any third party
                    sites or services.
                </Text>
                <Heading textAlign='left' mb='3' mt='7'>
                    Changes to this Privacy Policy
                </Heading>
                <Text>
                    We may update Our Privacy Policy from time to time. We will
                    notify You of any changes by posting the new Privacy Policy
                    on this page.
                </Text>
                <Text>
                    We will let You know via email and/or a prominent notice on
                    Our Service, prior to the change becoming effective and
                    update the &quot;Last updated&quot; date at the top of this
                    Privacy Policy.
                </Text>
                <Text>
                    You are advised to review this Privacy Policy periodically
                    for any changes. Changes to this Privacy Policy are
                    effective when they are posted on this page.
                </Text>
                <Heading textAlign='left' mb='3' mt='7'>
                    Contact Us
                </Heading>
                <Text>
                    If you have any questions about this Privacy Policy, You can
                    contact us:
                </Text>
                <UnorderedList>
                    <ListItem>
                        By email:{' '}
                        <Link
                            href='mailto:pacstall@protonmail.com'
                            color='pink.400'
                        >
                            pacstall@protonmail.com
                        </Link>
                    </ListItem>
                </UnorderedList>
            </Container>
        </PageAnimation>
    </>
)

export default PrivacyPolicy
