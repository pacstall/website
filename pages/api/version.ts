// Next.js API route support: https://nextjs.org/docs/api-routes/introduction
import type { NextApiRequest, NextApiResponse } from 'next'

export type LastestVersionDto = {
  version: string
}

const fetchVersion = async (): Promise<string> => {
  const res = await fetch('https://github.com/pacstall/pacstall/releases/latest')
  const allLines = await res.text();
  const content = allLines.split('\n').join('')
  const versionMatcher = /<a href="\/pacstall\/pacstall\/tree\/([1-9][0-9]*(\.[0-9]*)*)/
  const [_, version] = versionMatcher.exec(content) as string[]
  return version
}

let cachedVersion: string = ''

export default async function handler(
  req: NextApiRequest,
  res: NextApiResponse<LastestVersionDto>
) {
  if (!cachedVersion) {
    cachedVersion = await fetchVersion()
    setTimeout(() => cachedVersion = '', 5 * 60 * 1000)
  }

  res.status(200).json({ version: cachedVersion })
}
