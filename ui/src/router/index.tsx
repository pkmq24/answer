import React, { Suspense, lazy } from 'react';
import { RouteObject, createBrowserRouter } from 'react-router-dom';

import routeConfig, { RouteNode } from '@/router/route-config';
import RouteRules from '@/router/route-rules';

const routes: RouteObject[] = [];

const routeGen = (routeNodes: RouteNode[], root: RouteObject[]) => {
  routeNodes.forEach((rn) => {
    rn.page = rn.page.replace('pages/', '');
    /**
     * cannot use a fully dynamic import statement
     * ref: https://webpack.js.org/api/module-methods/#import-1
     */
    const Control = lazy(() => import(`@/pages/${rn.page}`));
    rn.element = (
      <Suspense>
        <Control />
      </Suspense>
    );
    root.push(rn);
    if (Array.isArray(rn.rules)) {
      const ruleFunc: Function[] = [];
      if (typeof rn.loader === 'function') {
        ruleFunc.push(rn.loader);
      }
      rn.rules.forEach((ruleKey) => {
        const func = RouteRules[ruleKey];
        if (typeof func === 'function') {
          ruleFunc.push(func);
        }
      });
      rn.loader = ({ params }) => {
        ruleFunc.forEach((func) => {
          func(params);
        });
      };
    }
    const children = Array.isArray(rn.children) ? rn.children : null;
    if (children) {
      rn.children = [];
      routeGen(children, rn.children);
    }
  });
};

routeGen(routeConfig, routes);

const router = createBrowserRouter(routes);
export default router;
