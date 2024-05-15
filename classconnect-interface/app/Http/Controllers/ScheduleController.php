<?php

namespace App\Http\Controllers;

use Illuminate\Http\Request;
use Illuminate\Support\Facades\Http;

class ScheduleController extends Controller
{
    /**
     * Показывает форму для загрузки расписания.
     *
     * @return \Illuminate\View\View
     */
    public function showForm()
    {
        return view('schedules.form');
    }

    /**
     * Обрабатывает загрузку расписания.
     *
     * @param  \Illuminate\Http\Request  $request
     * @return \Illuminate\Http\Response
     */
    public function upload(Request $request)
    {
        $token = session('access_token');

        // Проверяем наличие токена в сессии
        if (!$token) {
            // Если токен отсутствует, возвращаем сообщение об ошибке
            return response()->json(['error' => 'Unauthorized'], 401);
        }

        // Получаем текст из поля формы
        $text = $request->input('text');

        // Декодируем JSON для проверки корректности структуры
        $jsonData = json_decode($text, true);

        if (json_last_error() !== JSON_ERROR_NONE) {
            // Если произошла ошибка при декодировании JSON
            return response()->json(['error' => 'Invalid JSON format'], 400);
        }

        // Отправляем запрос на API для создания расписания
        $response = Http::withHeaders([
            'Authorization' => 'Bearer ' . $token,
            'Content-Type' => 'application/json',
        ])->post('http://172.26.0.4:8080/api/v1/schedules', $jsonData);

        // Проверяем успешность запроса
        if ($response->successful()) {
            // Если успешно, возвращаем ответ от API
            return $response->json();
        } else {
            // Если произошла ошибка, возвращаем ответ с соответствующим статусом
            return response()->json($response->json(), $response->status());
        }
    }
}
