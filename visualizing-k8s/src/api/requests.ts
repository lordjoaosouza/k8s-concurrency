import { cluster, single } from '@/services/request';

interface RequestProps {
  requests: number;
}

interface RequestResponse {
  singleRequestsTime: number;
  clusterRequestsTime: number;
}

export async function executeRequests({ requests }: RequestProps) {
  const singleResults = []
  const resultsCluster = []

  for (let i = 0; i < requests; i++) {
    singleResults.push(single.get('/picalc/100000'))
    resultsCluster.push(cluster.get('/picalc/100000'))
  }

  const singleRequestStartInstant = performance.now();
  await Promise.all(singleResults)
  const singleRequestEndInstant = performance.now();

  const singleRequestsTime = singleRequestEndInstant - singleRequestStartInstant;

  const clusterRequestStartInstant = performance.now();
  await Promise.all(resultsCluster)
  const clusterRequestEndInstant = performance.now();

  const clusterRequestsTime = clusterRequestEndInstant - clusterRequestStartInstant;

  return { singleRequestsTime, clusterRequestsTime } as RequestResponse;
}