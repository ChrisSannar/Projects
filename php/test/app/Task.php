<?php

namespace App;

use Illuminate\Database\Eloquent\Model;

class Task extends Model
{
    protected $guarded = [];

    public function user() {
        return $this->belongsTo(User::class);
    }

    public function setTask($newTask) {
        $this->attributes['text'] = $newTask['text'];
        $this->attributes['completed'] = $newTask['completed'];
    }
}
