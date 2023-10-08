import { StyledSidebar } from '@pm/components/sidebar/styles.ts';
import React, { Key, ReactNode, useEffect, useMemo } from 'react';
import { Menu, MenuProps } from 'antd';
import { AuditOutlined, UserSwitchOutlined } from '@ant-design/icons';
import { useLocation, useNavigate, useParams, useRoutes } from 'react-router-dom';

type MenuItem = Required<MenuProps>['items'][number];

const getMenuItem = (
  label: ReactNode,
  key: Key,
  icon?: ReactNode,
  children?: MenuItem[],
  type?: 'group',
  onClick?: () => void,
): MenuItem => {
  return {
    key,
    label,
    icon,
    children,
    type,
    onClick,
  } as MenuItem;
};

const rootMenuKeys: string[] = ['members', 'pa-pc-results'];

const Sidebar = () => {
  const [selectedKeys, setSelectedKeys] = React.useState<string[]>([]);
  const navigate = useNavigate();
  const path = useParams();
  const location = useLocation();

  const menuItems: MenuItem[] = useMemo(
    () => [
      getMenuItem('Members', 'members', <UserSwitchOutlined />, undefined, undefined, () =>
        navigate('/members'),
      ),
      getMenuItem('PA/PC Report', 'pa-pc-results', <AuditOutlined />, undefined, undefined, () =>
        navigate('/pa-pc-results'),
      ),
    ],
    [navigate],
  );

  const onOpenChange = (keys: string[]) => {
    const latestOpenKey = keys.find((key) => selectedKeys.indexOf(key) === -1);
    if (latestOpenKey && rootMenuKeys.indexOf(latestOpenKey!) === -1) {
      setSelectedKeys(keys);
    } else {
      setSelectedKeys(latestOpenKey ? [latestOpenKey] : []);
    }
  };

  useEffect(() => {
    const currentPath = location.pathname.split('/')[1];
    if (currentPath) {
      setSelectedKeys([currentPath]);
    }
  }, [location.pathname]);

  return (
    <StyledSidebar>
      <Menu
        mode='inline'
        selectedKeys={selectedKeys}
        onOpenChange={onOpenChange}
        items={menuItems}
      />
    </StyledSidebar>
  );
};

export default Sidebar;
