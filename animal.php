<?php

class Animal{
    
    public array $animals = [];
    
    public function __construct(array $data) 
    {
        $this->animals = $data;
    }

    public function index() 
    {
        $animals = $this->animals;

        foreach ($animals as $key =>$animal) {
            echo "index: $key, name: $animal" . PHP_EOL;
        }
    }

    public function store(string $data) 
    {
       $this->animals[] = $data;

        //array_push($this->animals, $data);
    }

    public function update($index, $data) 
    {
        $this->animals[$index] = $data;
    }

    public function destroy($index) 
    {
        unset($this->animals[$index]);
    }
}

$data = ['Kucing', 'Harimau', 'Jerapah', 'Ayam', 'Ikan'];

$animals = new Animal($data);

echo 'Menampilkan data animals' . PHP_EOL;
$animals->index();

echo PHP_EOL;

echo 'Menambahkan data baru' . PHP_EOL;
$animals->store('Kudanil');
$animals->store('Onta');
$animals->index();

echo PHP_EOL;

echo 'Mengupdate data: ' . PHP_EOL;
$animals->update(2, 'Kancil');
$animals->index();

echo PHP_EOL;

echo 'Menghapus data' . PHP_EOL;
$animals->destroy(3);
$animals->index();
?>