<!-- resources/views/schedules/form.blade.php -->

@extends('layouts.app')

@section('content')
    <div class="container">
        <div class="row justify-content-center">
            <div class="col-md-8">
                <div class="card">
                    <div class="card-header">Введите расписание</div>

                    <div class="card-body">
                        <form method="POST" action="{{ route('schedules.upload') }}">
                            @csrf

                            <!-- Поле для ввода текста -->
                            <div class="form-group">
                                <label for="text">json формат</label>
                                <textarea class="form-control" id="text" name="text" rows="5"></textarea>
                            </div>

                            <!-- Кнопка для отправки формы -->
                            <button type="submit" class="btn btn-primary">Отправить</button>
                        </form>
                    </div>
                </div>
            </div>
        </div>
    </div>
@endsection
