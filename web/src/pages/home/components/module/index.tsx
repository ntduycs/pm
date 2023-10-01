import { StyledModule } from '@pm/pages/home/components/module/styles.ts';
import { Card } from 'antd';
import { useNavigate } from 'react-router-dom';
import { ReactNode } from 'react';

type ModuleProps = {
  title: string;
  icon: ReactNode;
  path: string;
};

export const Module = ({ title, icon, path }: ModuleProps) => {
  const navigate = useNavigate();

  return (
    <StyledModule>
      <Card
        hoverable={true}
        cover={icon}
        onClick={() => navigate(path)}
      >
        <Card.Meta
          title={title}
          className='module-title'
        />
      </Card>
    </StyledModule>
  );
};
