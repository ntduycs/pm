import {
  EmptyResponse,
  ListPaPcResultsRequest,
  ListPaPcResultsResponse,
  UpsertPaPcResultRequest,
} from '@pm/models';
import { axios } from '@pm/services/index.ts';

export const listPaPcResultsAPI = async (
  req: ListPaPcResultsRequest,
): Promise<ListPaPcResultsResponse> => {
  const response = await axios.get(`pa-pc-results`, {
    params: req,
  });

  return response.data;
};

export const upsertPaPcResultAPI = async (
  paPcResult: UpsertPaPcResultRequest,
): Promise<EmptyResponse> => {
  const response = await axios.post(`pa-pc-results`, paPcResult, {
    headers: {
      'Content-Type': 'application/json',
      Accept: 'application/json',
    },
    responseType: 'json',
  });

  if (response.status > 299) {
    throw new Error(response.data.message);
  }

  return response.data;
};
