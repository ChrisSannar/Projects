@extends('layouts.app')

@section('content')
<div class="card">
    <div class="card-header">
        <h2>Tasks</h2>
        <span>You have {{ $tasks->count() }} tasks</span>
    </div>
    <div class="card-body">
    <ul>
        @foreach($tasks as $task)
            <div class="row">
                <li>{{ $task->text }}</li>
                <input type="checkbox" {{ $task->completed ? 'checked="checked"' : '' }} name="done" id="done" >
            </div>
        @endforeach
    </ul>
    <div class="row">
        <a href="/tasks/create">Add Task</a>
    </div>
    <div class="row">
        <a href="/profile">Go Back</a>
    </div>
    </div>
</div>
@endsection
