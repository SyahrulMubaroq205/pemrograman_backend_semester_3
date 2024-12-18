const persiapan = () => {
    setTimeout(function() {
    console.log("Mempersiapkan Bahan...");
    }, 3000);
};
const rebusAir = () => {
    setTimeout(function() {
    console.log("Merebus Air...");
    }, 5000);
};

const masak = () => {
    setTimeout(function() {
    console.log("Memasak Mie...");
    }, 7000);
};

const selesai = () => {
    setTimeout(function() {
    console.log("Selesai");
    }, 9000)
}

const main = () => {
    persiapan();
    rebusAir();
    masak();
    selesai();
};

main();