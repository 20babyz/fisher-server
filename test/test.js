import http from 'k6/http';
import {check, sleep} from 'k6';

export let options = {
    stages: [
        { duration: '10s', target: 10 }, // 10초 동안 10명의 동시 사용자
        { duration: '30s', target: 50 }, // 30초 동안 50명의 동시 사용자 증가
        { duration: '50s', target: 100 }, // 50초 동안 100명의 동시 사용자 증가
        { duration: '30s', target: 10000 }, // 30초 동안 10000명의 동시 사용자 증가
        { duration: '10s', target: 0 },  // 10초 동안 점진적 종료
    ],
};

export default function () {
    let res1 = http.get('http://localhost:8080/');
    check(res1, { 'status is 200': (r) => r.status === 200 });

    let payload = JSON.stringify({ name: "John Doe", age: 30 });
    let params = { headers: { 'Content-Type': 'application/json' } };
    let res2 = http.post('http://localhost:8080/post', payload, params);
    check(res2, { 'status is 200': (r) => r.status === 200 });

    sleep(1); // 1초 대기 후 반복
}