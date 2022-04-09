import { Box, chakra, useBreakpointValue } from "@chakra-ui/react";
import { FC } from "react";
import { Helmet } from "react-helmet";
import PageAnimation from "../components/animations/PageAnimation";

const ChatBubbleOption0Desktop: FC = () => (<>
    ╭───────────────────────────────────────────────────────────╮<br />
    │ Oh, you thought it was <strong>404</strong> but it was me, Pac the alpaca! │<br />
    ╰───────────────────────────────────────────────────────────╯
</>)

const ChatBubbleOption1Desktop: FC = () => (<>
    ╭─────────────────────────────────────────────────────────╮<br />
    │ Ooops, I can't find that page using my Alpaca skills :( │<br />
    ╰─────────────────────────────────────────────────────────╯
</>)

const ChatBubbleOption2Desktop: FC = () => (<>
    ╭────────────────────────────────────────────────────────────────╮<br />
    │ Hey, cheer up. Instead of finding the page, you found me, Pac! │<br />
    ╰────────────────────────────────────────────────────────────────╯
</>)

const ChatBubbleOption3Desktop: FC = () => (<>
    ╭─────────────────────────────────────────────────────────────────────────╮<br />
    │ Have you seen my cousin Stall? They were supposed to be on this page... │<br />
    ╰─────────────────────────────────────────────────────────────────────────╯
</>)

const ChatBubbleOption4Desktop: FC = () => (<>
    ╭────────────────────────────────────────────────╮<br />
    │ Alpaca your bags, 'cause there is nothing here │<br />
    ╰────────────────────────────────────────────────╯
</>)

const ChatBubbleOption0Mobile: FC = () => (<>
    <br />╭───────────────────────────╮
    <br />│ Oh, you thought it was <strong>404</strong>      │
    <br />│ but it was me, Pac the alpaca!  │
    <br />╰───────────────────────────╯
</>)

const ChatBubbleOption1Mobile: FC = () => (<>
    <br />╭──────────────────────────╮
    <br />│ Ooops, I can't find that page  |
    <br />| using my Alpaca skills :(      │
    <br />╰──────────────────────────╯
</>)

const ChatBubbleOption2Mobile: FC = () => (<>
    <br />╭──────────────────────────────╮
    <br />│ Hey, cheer up. Instead of finding   |
    <br />| the page, you found me, Pac!        │
    <br />╰──────────────────────────────╯
</>)

const ChatBubbleOption3Mobile: FC = () => (<>
    <br />╭───────────────────────────╮
    <br />│ Have you seen my cousin Stall?  |
    <br />| They were supposed to be        |
    <br />| on this page...                 │
    <br />╰───────────────────────────╯
</>)

const ChatBubbleOption4Mobile: FC = () => (<>
    <br />╭─────────────────────────╮
    <br />│ Alpaca your bags,            |
    <br />|'cause there is nothing here  │
    <br />╰─────────────────────────╯
</>)

const ChatBubble: FC = () => {
    const anyOf = (items: any[]) => items[Math.floor(Math.random() * items.length)]
    const desktopOptions = [<ChatBubbleOption0Desktop />, <ChatBubbleOption1Desktop />, <ChatBubbleOption2Desktop />, <ChatBubbleOption3Desktop />, <ChatBubbleOption4Desktop />]
    const mobileOptions = [<ChatBubbleOption0Mobile />, <ChatBubbleOption1Mobile />, <ChatBubbleOption2Mobile />, <ChatBubbleOption3Mobile />, <ChatBubbleOption4Mobile />]

    const options = useBreakpointValue({ base: mobileOptions, md: desktopOptions })

    return <>{anyOf(options)}</>
}

const NotFound: FC = () => (
    <>
        <Helmet>
            <title>Not Found - Pacstall</title>
        </Helmet>
        <PageAnimation>
            <Box textAlign='center' mt='20vh'>
                <chakra.pre style={{ fontFamily: 'JetBrains Mono' }} fontSize={useBreakpointValue({ base: 'sm', md: 'md' })}>
                    <code>
                        <ChatBubble />
                        <br />\     ∩~-~∩
                        <br />  \   ξ •×• ξ
                        <br />      ξ　~  ξ
                        <br />      ξ　   ξ
                        <br />              ξ　   “~~~~~~〇
                        <br />             ξ　          ξ
                        <br />              ξ ξ ξ~~~~~ξξ ξ
                        <br />              ξ_ξ ξ_ξ ξ_ξξ_ξ
                        <br />             404 Page Not Found
                    </code>
                </chakra.pre>
            </Box>
        </PageAnimation>
    </>
)

export default NotFound