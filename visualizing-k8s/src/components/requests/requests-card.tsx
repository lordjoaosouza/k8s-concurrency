import Image from 'next/image'

import { Card, CardContent, CardDescription, CardHeader, CardTitle } from '@/components/ui/card'
import { Label } from '@/components/ui/label'
import { Progress } from '@/components/ui/progress'

interface ResultsCardProps {
  title: string
  description: string
  icon: string
  progressValue: number
}

export default function RequestsCard({
  title,
  description,
  icon,
  progressValue,
}: ResultsCardProps) {
  return (
    <div>
      <Card className={'h-96 w-96'}>
        <CardHeader>
          <CardTitle className={'flex items-center gap-2 text-lg'}>
            <Image src={`${icon}.svg`} alt={icon} width={40} height={40} /> {title}
          </CardTitle>
          <CardDescription>{description}</CardDescription>
        </CardHeader>
        <CardContent className={'flex h-full flex-col items-center justify-center'}>
          <div className={'h-1/2 w-full'}>
            <Label>Requesting...</Label>
            <Progress value={progressValue} />
          </div>
        </CardContent>
      </Card>
    </div>
  )
}
