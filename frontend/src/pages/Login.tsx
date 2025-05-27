import {createSignal, JSX, Show} from "solid-js";
import {A, useNavigate} from "@solidjs/router";
import {Button, IconButton, TextField} from "@suid/material";
import login, {LoginForm} from "../api/user.ts";
import {Visibility, VisibilityOff} from "@suid/icons-material";

interface LoginStatus {
    success: boolean,
    msg: string
}

export default function Login(): JSX.Element {
    const [formData, setFormData] = createSignal<LoginForm>({
        username: '',
        password: ''
    })
    const [loading, setLoading] = createSignal<boolean>(false)
    const [show, setShow] = createSignal<boolean>(false)
    const [msg, setMsg] = createSignal<LoginStatus>({
        success: false,
        msg: ''
    })
    const handleChange = (e: Event) => {
        const target = e.currentTarget as HTMLInputElement;
        setFormData(prev => ({
            ...prev,
            [target.name]: target.value
        }));
    }
    const nav = useNavigate();
    const [showPassword, setShowPassword] = createSignal<boolean>(false);

    function handleClickShowPassword() {
        setShowPassword(!showPassword())
    }

    async function handleSubmit(e: Event) {
        e.preventDefault();
        setLoading(true)
        const res = await login(formData())
        if (res.data.msg === "success") {
            setMsg({
                success: true,
                msg: "登录成功"
            })
            setShow(true)
            setTimeout(() => {
                setShow(false)
                nav("/home")
            }, 1000)
        } else {
            setMsg({
                success: false,
                msg: res.data.msg
            })
            setShow(true)
            setTimeout(() => {
                setShow(false)
            }, 3000)
        }

        setLoading(false)
    }

    return (
        <div class="min-h-screen flex items-center justify-center bg-gray-100 dark:bg-gray-900 p-4">
            <div
                class="flex flex-col gap-6 w-full max-w-md bg-white dark:bg-gray-800 rounded-xl shadow-lg p-6 md:p-8 mx-auto">
                <h2 class="text-2xl font-bold mb-6 text-center text-gray-800 dark:text-gray-100">
                    Admin 登录
                </h2>
                <form onSubmit={handleSubmit} class="flex flex-col gap-4">
                    <TextField
                        label="用户名" variant="outlined" fullWidth required
                        autoFocus={true}
                        autoComplete={"username"}
                        value={formData().username}
                        onChange={handleChange}
                        name={"username"}
                    />

                    <TextField
                        label="密码" variant="outlined" fullWidth required
                        autoFocus={true}
                        autoComplete={"current-password"}
                        value={formData().password}
                        onChange={handleChange}
                        name={"password"}
                        InputProps={{
                            endAdornment: (
                                <IconButton
                                    aria-label="toggle password visibility"
                                    onClick={handleClickShowPassword}
                                    edge="end"
                                >
                                    {showPassword() ? <VisibilityOff/> : <Visibility/>}
                                </IconButton>
                            ),
                        }}
                        type={showPassword() ? 'text' : 'password'}
                    />
                    <Button
                        disabled={loading()}
                        type="submit" variant="contained" color="primary" class="mt-2 pt-10">
                        登录
                    </Button>

                </form>
                <div class="mt-6 text-sm text-center text-gray-600 dark:text-gray-400">
                    还没有账户吗？
                    <A href="/user/register"
                       replace={true}
                       class="ml-1 text-blue-500 hover:underline">
                        在此注册
                    </A>
                </div>
                <Show when={show()} fallback={<></>}>
                   <span
                       class={`text-center text-white px-4 py-2
                       rounded-md shadow-md text-sm transition-colors duration-300
                        ${msg().success ? 'bg-green-500' : 'bg-red-500'}`}
                   >
                        {msg().msg}
                    </span>
                </Show>
            </div>
        </div>
    )
}