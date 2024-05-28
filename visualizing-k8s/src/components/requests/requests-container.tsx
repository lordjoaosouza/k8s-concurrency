import { Repeat } from 'lucide-react'
import Link from 'next/link'
import { useEffect } from 'react'

import RequestsCard from '@/components/requests/requests-card'
import { Button } from '@/components/ui/button'
import { cluster, single } from '@/services/request'

interface RequestsContainerProps {
  requests: number
  setIsRunning: (isRunning: boolean) => void
}

export default function RequestsContainer({ requests, setIsRunning }: RequestsContainerProps) {
  useEffect(() => {
    const handleRequests = async () => {
      const singleResults = []
      const resultsCluster = []
      for (let i = 0; i < requests; i++) {
        singleResults.push(single.get('/picalc/100000'))

        resultsCluster.push(cluster.get('/picalc/100000'))
      }

      try {
        const singleResponses = await Promise.all(singleResults)
        const clusterResponses = await Promise.all(resultsCluster)

        console.log(singleResponses)
        console.log(clusterResponses)
      } catch (error) {
        console.error('Error trying to request', error)
      }
    }

    handleRequests().then(() => {})
  }, [])

  return (
    <div className={'flex flex-col items-center gap-4'}>
      <div className={'flex w-full items-center justify-between'}>
        <h1>Stress testing with {requests} requests</h1>

        <Link href={'/'}>
          <Button
            className={'gap-1'}
            onClick={() => {
              setIsRunning(false)
            }}
          >
            New test <Repeat className={'size-4'} />
          </Button>
        </Link>
      </div>

      <div className={'flex gap-4'}>
        <RequestsCard
          title={'Single Docker container'}
          description={'With no concurrency'}
          icon={'docker'}
          progressValue={1}
        />

        <RequestsCard
          title={'Kubernetes cluster'}
          description={'With concurrency'}
          icon={'k8s'}
          progressValue={20}
        />
      </div>
    </div>
  )
}
