import { Container, Heading, Image, Link, Stack, Text } from '@chakra-ui/react'
import { FC } from 'react'
import Card from '../components/Card'
import Navigation from '../components/Navigation'

// @ts-ignore:next-line
import PacstallLogo from '../../public/pacstall.svg'
import { Helmet } from 'react-helmet'


const Home: FC = () => {
	return (<>
		<Helmet>
			<title>Pacstall - The AUR for Ubuntu</title>
		</Helmet>
		<Navigation />
		<Container maxW='900px'>

			<Heading size='2xl' textAlign='center' p='7'>Pacstall ~ The AUR for Ubuntu</Heading>

			<Stack maxW='lg' margin='auto'>
				<Card title='Why is this any different than any other package manager?'>
					<Text>
						Pacstall uses the stable base of Ubuntu but allows you to use bleeding edge software with little to no compromises, so you don't have to worry about security patches or new features.
					</Text>
				</Card>
			</Stack>

			<Stack maxW='lg' margin='auto'>
				<Card title='How does it work then?'>
					<Text>
						Pacstall takes in files known as <Link color='pink.400' href="https://github.com/pacstall/pacstall/wiki/Pacscript-101">pacscripts</Link> (similar to PKGBUILD's) that contain the necessary contents to build packages, and builds them into
						executables on your system.
					</Text>
				</Card>
			</Stack>

			<Image
				src={PacstallLogo}
				width="200px"
				height="200px"
				alt="Pacstall logo"
				m='auto'
				mt='6'
				loading="lazy" />
		</Container>
	</>
	)
}

export default Home
