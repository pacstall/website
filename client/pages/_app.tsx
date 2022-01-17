import '../styles/globals.css'
import '../styles/dracula.css'
import type { AppProps } from 'next/app'
import Metadata from '../components/Head'

function MyApp({ Component, pageProps }: AppProps) {
  return <>
    <Metadata />
    <Component {...pageProps} />
  </>
}

export default MyApp
