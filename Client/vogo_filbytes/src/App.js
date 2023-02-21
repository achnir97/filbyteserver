
import './App.css';
import  Header from  './Vogo'
import FourColumns from './Vogo_1';
import TwoColumns from './Vogo_2';
import Frpinv from './Vogo_3';
import Vogo_4 from './Vogo_4';
import Tabs from './Browser';
function App() {

  return (
    <div className="App" >
      
       <TwoColumns column1="VOGO DIGITAL LAB" column2="ID"/>
       <FourColumns column1="계약 체결일" column3="FIL 지급 개시일"/>
       <FourColumns column1="게야된 FRP 유효채글파워" column3="투자지원금"/>
       <FourColumns column1="현재 FRP 유효채글파워" column3="계약일 매입가능수량"/>
       <FourColumns column1="서약금" column3="담보 FIL매입수량"/>
       <TwoColumns column1="주소" column2="서울시 영동포구 여의나루로 12, 2 동 706 호 " />
       <FourColumns column1="연락처" column3="이메일"/>
       <FourColumns column1="거래은행" column3="게좌번호"/>
       <TwoColumns column1="FIL 지갑주소" column2="f1hgb4b5qf7aliphgi2emxoa6olwvpxb3xvuwdply"/>
      <br/>

       <TwoColumns column1="KSL 실시간 현황" column2="날자 / FIL 가격"/>
       <FourColumns column1="당월 보상 FIL" column3="당원 staking FIL "/>
       <FourColumns column1="당월 FRP 증가" column3="당월 지급 FIL"/>
       <FourColumns column1="총 STAKING 수량(FIL) 및
       총 우효보상 파워" column2="총 지급수량(FIL) 및 현재가치(원)" column3= "총 보상 수량(FIL) 및 현재가치(원)" column4="목표수일류 및 현재 달설율(%)"/>
       <FourColumns/>
       <FourColumns/>
   
       <FourColumns column1='총 STAKING 수량 및  유효보상 파워(TB)' column2='총 지급수량(FIL) 및 현재 가치' column3='총 보상수량 및 현재가치'
       column4='목펴수일률 및 현제 달설률'/>
       <FourColumns/>
       <FourColumns/>
       <Vogo_4/>
       
       
    </div>
  );
}

export default App;
