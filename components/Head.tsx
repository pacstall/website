import Head from "next/head";
import { FC } from "react";

const Metadata: FC = () => (
    <Head>
        <title>Pacstall - The AUR for Ubuntu</title>
        <meta name="title" content="Pacstall - The AUR for Ubuntu" />
        <meta name="description" content="Pacstall automates downloading source packages, installing dependencies, and installing, in Ubuntu" />
        <meta property="og:type" content="website" />
        <meta property="og:url" content="https://pacstall.dev" />
        <meta property="og:title" content="Pacstall - The AUR for Ubuntu" />
        <meta property="og:description" content="Pacstall automates downloading source packages, installing dependencies, and installing, in Ubuntu" />
        <meta property="og:image" content="pacstall.svg" />

        <meta property="twitter:card" content="summary_large_image" />
        <meta property="twitter:url" content="https://pacstall.dev" />
        <meta property="twitter:title" content="Pacstall - The AUR for Ubuntu" />
        <meta property="twitter:description" content="Pacstall automates downloading source packages, installing dependencies, and installing, in Ubuntu" />
        <meta property="twitter:image" content="pacstall.svg" />
        <meta name="apple-mobile-web-app-title" content="Pacstall" />
        <meta name="apple-mobile-web-app-capable" content="yes" />

        <meta name="viewport" content="width=device-width, initial-scale=1" />
        <link rel="shortcut icon" type="image/jpg" href="favicon.ico" />
        <link rel="image_src" href="pacstall.svg" />
        <link rel="apple-touch-icon" href="apple-touch-icon.png" />
    </Head>
)

export default Metadata