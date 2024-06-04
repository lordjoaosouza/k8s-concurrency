import Image from 'next/image'

import { Card, CardContent, CardDescription, CardHeader, CardTitle } from '@/components/ui/card'
import {
  Table,
  TableBody,
  TableCaption,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from '@/components/ui/table'

interface ResultsCardProps {
  title: string
  description: string
  icon: string
  requestAmount: number
  timeInMilliseconds: number
  hits: number
  fails: number
}

export default function RequestsCard({
  title,
  description,
  icon,
  requestAmount,
  timeInMilliseconds,
  hits,
  fails,
}: ResultsCardProps) {
  return (
    <div>
      <Card className={'h-[450px] w-96'}>
        <CardHeader>
          <CardTitle className={'flex items-center gap-2 text-lg'}>
            <Image src={`${icon}.svg`} alt={icon} width={40} height={40} /> {title}
          </CardTitle>
          <CardDescription>{description}</CardDescription>
        </CardHeader>
        <CardContent className={'flex h-[calc(450px-112px)] flex-col items-center justify-center'}>
          <Table>
            <TableCaption>A list of stress test results</TableCaption>

            <TableHeader>
              <TableRow>
                <TableHead className='w-[200px]'>Metric</TableHead>
                <TableHead className='text-right'>Amount</TableHead>
              </TableRow>
            </TableHeader>

            <TableBody>
              <TableRow>
                <TableCell className='font-medium'>Requests amount</TableCell>
                <TableCell className='text-right'>{requestAmount}</TableCell>
              </TableRow>

              <TableRow>
                <TableCell className='font-medium'>Total request time</TableCell>
                <TableCell className='text-right'>{timeInMilliseconds}ms</TableCell>
              </TableRow>

              <TableRow>
                <TableCell className='font-medium'>Hits</TableCell>
                <TableCell className='text-right'>{hits}</TableCell>
              </TableRow>

              <TableRow>
                <TableCell className='font-medium'>Fails</TableCell>
                <TableCell className='text-right'>{fails}</TableCell>
              </TableRow>
            </TableBody>
          </Table>
        </CardContent>
      </Card>
    </div>
  )
}
