import http from 'k6/http';
import { check, sleep } from 'k6';

export let options = {
  vus: 5000,
  duration: '30s',
};

export default function () {
  // Mendapatkan timestamp sebelum mengirim permintaan
  let startTime = new Date().getTime();


  let res = http.get('https://gin-production-3247.up.railway.app/');

  // Memeriksa apakah respons memiliki status 200
  check(res, {
    'status is 200': (r) => r.status === 200,
  });

  // Mendapatkan timestamp setelah menerima respons
  let endTime = new Date().getTime();

  // Menghitung waktu respons dalam milidetik (ms)
  let responseTime = endTime - startTime;
  // console.log(`Response time: ${responseTime} ms`);

  // Menunggu 1 detik sebelum melakukan permintaan berikutnya
  sleep(1);
}