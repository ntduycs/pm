import { PageResponse } from '@pm/models';
import { ReactNode } from 'react';
import { StyledTableFooter } from '@pm/components/table-footer/styles.ts';

type TableFooterProps = {
  data?: PageResponse<unknown>;
};

const TableFooter = ({ data: res }: TableFooterProps): ReactNode => {
  if (!res) return null;

  if (res.total === 0) {
    return <StyledTableFooter>No data found.</StyledTableFooter>;
  }

  const from = (res?.page - 1) * res?.size + 1;
  const to = from + res?.items.length - 1;

  return (
    <StyledTableFooter>
      Showing {res?.count} ({from} - {to}) of {res.total} items.
    </StyledTableFooter>
  );
};

export default TableFooter;
