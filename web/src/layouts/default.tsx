import { StyledDefaultLayout } from '@pm/layouts/styles.ts';
import { Outlet } from 'react-router-dom';
import { Sidebar } from '@pm/components';
import { Col, Row } from 'antd';

export const DefaultLayout = () => {
  return (
    <StyledDefaultLayout>
      <Row>
        <Col
          xs={24}
          sm={24}
          md={24}
          lg={4}
          className='aside'
        >
          <Sidebar />
        </Col>
        <Col
          xs={24}
          sm={24}
          md={24}
          lg={20}
          className='content'
        >
          <Outlet />
        </Col>
      </Row>
    </StyledDefaultLayout>
  );
};
