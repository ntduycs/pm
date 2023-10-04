import { StyledHome } from '@pm/pages/home/styles.ts';
import { Col, Row } from 'antd';
import { AuditOutlined, UserSwitchOutlined } from '@ant-design/icons';
import { Module } from '@pm/pages/home/components';

export const Home = () => {
  return (
    <StyledHome>
      <Row>
        <Col span={6}>
          <Module
            title={'Member Management'}
            icon={<UserSwitchOutlined className='module-icon' />}
            path={'/members'}
          />
        </Col>
        <Col span={6}>
          <Module
            title={'PA/PC Report'}
            icon={<AuditOutlined className='module-icon' />}
            path={'/pa-pc'}
          />
        </Col>
      </Row>
    </StyledHome>
  );
};
