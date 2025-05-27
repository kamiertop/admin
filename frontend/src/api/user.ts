import {POST} from "./request.ts";

export interface LoginForm {
    username: string
    password: string
}


export default async function login(data: LoginForm) {
    return await POST<{ msg: string }>('/user/login', data)
}