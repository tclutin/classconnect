<?php

namespace App\Http\Controllers;

use Illuminate\Http\Request;
use Illuminate\Support\Facades\Http;
use Illuminate\Support\Facades\Redirect;
use Illuminate\Support\Facades\Session;

class AuthController extends Controller
{
    /**
     * Показывает форму регистрации пользователя.
     *
     * @return \Illuminate\View\View
     */
    public function showSignupForm()
    {
        return view('signup');
    }

    /**
     * Регистрирует нового пользователя.
     *
     * @param  \Illuminate\Http\Request  $request
     * @return \Illuminate\Http\Response
     */
    public function signup(Request $request)
    {
        // Отправляем запрос на API для регистрации пользователя
        $response = Http::post('http://172.26.0.4:8080/api/v1/auth/signup', [
            'username' => $request->username,
            'email' => $request->email,
            'password' => $request->password,
        ]);

        // Проверяем статус ответа
        if ($response->status() === 201) {
            $request->session()->put('access_token', $response->json()['access_token']);
            // Если регистрация прошла успешно, автоматически аутентифицируем пользователя
            // и перенаправляем на страницу /me
            return Redirect::route('me')->with('success', 'Registration successful');
        } else {
            // Если произошла ошибка, возвращаем ответ с соответствующим статусом
            return response()->json($response->json(), $response->status());
        }
    }

    /**
     * Показывает форму входа пользователя.
     *
     * @return \Illuminate\View\View
     */
    public function showLoginForm()
    {
        return view('login');
    }

    /**
     * Аутентифицирует пользователя.
     *
     * @param  \Illuminate\Http\Request  $request
     * @return \Illuminate\Http\Response
     */
    public function login(Request $request)
    {
        // Отправляем запрос на API для аутентификации пользователя
        $response = Http::post('http://172.26.0.4:8080/api/v1/auth/login', [
            'username' => $request->username,
            'password' => $request->password,
        ]);

        // Проверяем успешность запроса
        if ($response->successful()) {
            // Если успешно, сохраняем токен в сессии
            $request->session()->put('access_token', $response->json()['access_token']);
            // Редирект на страницу /me
            return Redirect::route('me');
        } else {
            // Если запрос не удался, возвращаем ошибку
            return response()->json(['error' => 'Unauthorized'], 401);
        }
    }


    /**
     * Показывает информацию о пользователе.
     *
     * @param  \Illuminate\Http\Request  $request
     * @return \Illuminate\View\View|\Illuminate\Http\RedirectResponse
     */
    public function showUserInfo(Request $request)
    {
        // Получаем токен из сессии
        $token = $request->session()->get('access_token');

        // Проверяем наличие токена в сессии
        if (!$token) {
            // Если токен отсутствует, перенаправляем на страницу входа
            return redirect()->route('login')->with('error', 'You are not authenticated.');
        }

        // Получаем информацию о пользователе из API с использованием токена
        $response = Http::withHeaders([
            'Authorization' => 'Bearer ' . $token
        ])->get('http://172.26.0.4:8080/api/v1/auth/me');

        // Проверяем, успешно ли прошел запрос к API
        if ($response->successful()) {
            // Если успешно, возвращаем представление с данными пользователя
            $user = $response->json();
            return view('me', compact('user'));
        } else {
            // Если запрос не удался, перенаправляем на страницу входа с сообщением об ошибке
            return redirect()->route('login')->with('error', 'You are not authenticated.');
        }
    }

    public function logout()
    {
        // Очищаем сессию
        Session::flush();

        // Перенаправляем пользователя на страницу логина с сообщением об успешном выходе
        return Redirect::route('login')->with('success', 'You have been logged out successfully.');
    }
}
