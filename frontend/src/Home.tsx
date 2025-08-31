import type {JSX} from "solid-js";
import {type RouteProps} from "@solidjs/router";


// 布局组件, 在组件中嵌套路由, 实现header和footer不变, 切换路由时, 内容区域自动切换响应组件
export function AppLayout(props: RouteProps<any>): JSX.Element{
    return (
        <>
            <header>
                <nav></nav>
            </header>
            <main>
                {props.children}
            </main>
            <footer>
            </footer>
        </>
    )
}

export default function Home(props: RouteProps<any>): JSX.Element {
    return (
        <div>
            {props.children}
        </div>
    )
}