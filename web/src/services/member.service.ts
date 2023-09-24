import { EmptyResponse, ListMembersResponse, UpsertMemberRequest } from '@pm/models';
import { axios } from '@pm/services/index.ts';

export const listMembersAPI = async (): Promise<ListMembersResponse> => {
  // using Axios
  const response = await axios.get(`members`);

  return response.data;
};

export const upsertMemberAPI = async (member: UpsertMemberRequest): Promise<EmptyResponse> => {
  const response = await axios.post(`members`, member, {
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
