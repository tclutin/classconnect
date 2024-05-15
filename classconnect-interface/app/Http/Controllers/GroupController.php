<?php

// GroupController.php

namespace App\Http\Controllers;

use Illuminate\Http\Request;
use Illuminate\Support\Facades\Http;
use Illuminate\Support\Collection;
use Illuminate\Support\Facades\Redirect;

class GroupController extends Controller
{
    /**
     * Показывает список всех групп.
     *
     * @return \Illuminate\View\View
     */
    public function index()
    {
        // Получаем токен из сессии
        $token = session('access_token');

        // Проверяем наличие токена в сессии
        if (!$token) {
            // Если токен отсутствует, возвращаем сообщение об ошибке
            return response()->json(['error' => 'Unauthorized'], 401);
        }

        // Отправляем запрос на API для получения списка групп с использованием токена
        $response = Http::withHeaders([
            'Authorization' => 'Bearer ' . $token
        ])->get('http://172.26.0.4:8080/api/v1/groups');

        // Проверяем статус ответа
        if ($response->successful()) {
            // Если запрос прошел успешно, преобразуем данные в коллекцию объектов
            $groupsData = collect($response->json());

            // Преобразуем массив данных в коллекцию объектов
            $groups = $groupsData->map(function ($group) {
                return (object) $group;
            });

            // Возвращаем представление для отображения списка групп
            return view('groups.index', compact('groups'));
        } else {
            // Если произошла ошибка, возвращаем ответ с соответствующим статусом
            return response()->json($response->json(), $response->status());
        }
    }

    /**
    * Показывает форму для создания группы.
    *
    * @return \Illuminate\Http\Response
    */
    public function showCreateForm()
    {
        return view('groups.create');
    }

    /**
     * Создает новую группу.
     *
     * @param  \Illuminate\Http\Request  $request
     * @return \Illuminate\Http\Response
     */
    public function create(Request $request)
    {
        // Получаем токен из сессии
        $token = session('access_token');

        // Проверяем наличие токена в сессии
        if (!$token) {
            // Если токен отсутствует, возвращаем сообщение об ошибке
            return response()->json(['error' => 'Unauthorized'], 401);
        }

        // Отправляем запрос на API для создания группы
        $response = Http::withHeaders([
            'Authorization' => 'Bearer ' . $token
        ])->post('http://172.26.0.4:8080/api/v1/groups', [
            'name' => $request->name,
        ]);

        // Проверяем статус ответа
        if ($response->status() === 201) {
            // Если создание группы прошло успешно, перенаправляем пользователя на страницу со списком всех групп
            return Redirect::route('groups.index')->with('success', 'Group created successfully');
        } else {
            // Если произошла ошибка, возвращаем ответ с соответствующим статусом
            return response()->json($response->json(), $response->status());
        }
    }
}
