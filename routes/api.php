<?php

use App\Http\Controllers\StudentController;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\Route;

Route::get('/student', [StudentController::class, 'index']);
Route::post('/student', [StudentController::class, 'store']);
Route::patch('/student/{student}', [StudentController::class, 'update']);
Route::delete('/student/{student}', [StudentController::class, 'destroy']);
