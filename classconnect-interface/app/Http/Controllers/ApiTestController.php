<?php

namespace App\Http\Controllers;

use Illuminate\Http\Client\RequestException;
use Illuminate\Support\Facades\Http;

class ApiTestController extends Controller
{
    public function testApi()
    {
        try {
            $response = Http::get('http://172.26.0.4:8080/api/v1/auth/me');
            $data = $response->json();

            // Обработка успешного ответа от API
            return response()->json($data);
        } catch (RequestException $e) {
            // Обработка ошибки при запросе к API
            return response()->json(['error' => $e->getMessage()], 500);
        }
    }
}
