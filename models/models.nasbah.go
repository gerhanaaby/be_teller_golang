/**
 * @author [Fajar Dwi Nur Racmadi]
 * @email [fajar.d.rachmadi@banksinarmas.com]
 * @create date 2023-02-28
 * @modify date 2023-02-28
 * @desc [description]
 */
package models

import "gorm.io/gorm"

 type Nasabah struct {
	gorm.Model
	 
	 ID           		    int64   `json:"userid" gorm:"type:serial;primary_key"`
	 CIF                    string  `gorm:" column:CIF;type:varchar(256); null"`
	 Nama                   string  `gorm:" column:Nama;type:varchar(200); null"`
	 Tipe                   string  `gorm:" column:Tipe;type:varchar(200); null"`
	 KodeCabang             string  `gorm:" column:KodeCabang;type:varchar(200); null"`
	 NamaIbu                string  `gorm:" column:NamaIbu;type:varchar(200); null"`
	 TelpRumah              string  `gorm:" column:TelpRumah;type:varchar(200); null"`
	 Alamat                 string  `gorm:" column:Alamat;type:varchar(200); null"`
	 Email                  string  `gorm:" column:Email;type:varchar(200); null"`
	 CityTown               string  `gorm:" column:CityTown;type:varchar(200); null"`
	 Postcode               string  `gorm:" column:Postcode;type:varchar(200); null"`
	 Income                 float32 `gorm:" column:Income;type:float; null"`
	 Purpose                string  `gorm:" column:Purpose;type:varchar(200); null"`
	 Source                 string  `gorm:" column:Source;type:varchar(200); null"`
	 IDType                 string  `gorm:" column:IDType;type:varchar(200); null"`
	 Ponsel                 string  `gorm:" column:Ponsel;type:varchar(200); null"`
	 Type                   string  `gorm:" column:Type;type:varchar(200); null"`
	 Createdby              string  `gorm:" column:Createdby;type:varchar(200); null"`
	 IdNumber               string  `gorm:" column:IdNumber;type:varchar(100); null"`
	 Mnemonics              string  `gorm:" column:mnemonics;type:varchar(200); null"`
	 DocumentFlag           string  `gorm:" column:documentFlag;type:varchar(200); null"`
	 CountryCode            string  `gorm:" column:countryCode;type:varchar(200); null"`
	 Nationality            string  `gorm:" column:nationality;type:varchar(200); null"`
	 Email2                 string  `gorm:" column:email2;type:varchar(200); null"`
	 Email3                 string  `gorm:" column:email3;type:varchar(200); null"`
	 Email4                 string  `gorm:" column:email4;type:varchar(200); null"`
	 Target                 string  `gorm:" column:target;type:varchar(200); null"`
	 BiOwnership            string  `gorm:" column:biOwnership;type:varchar(200); null"`
	 OfficerCode            string  `gorm:" column:officerCode;type:varchar(200); null"`
	 Info1                  string  `gorm:" column:info1;type:varchar(200); null"`
	 Info2                  string  `gorm:" column:info2;type:varchar(200); null"`
	 Pillar                 string  `gorm:" column:pillar;type:varchar(200); null"`
	 Village                string  `gorm:" column:village;type:varchar(200); null"`
	 Subdistrict            string  `gorm:" column:subdistrict;type:varchar(200); null"`
	 Province               string  `gorm:" column:province;type:varchar(200); null"`
	 Occupation             string  `gorm:" column:occupation;type:varchar(200); null"`
	 BankRelation           string  `gorm:" column:bankRelation;type:varchar(200); null"`
	 BmpkViolation          string  `gorm:" column:bmpkViolation;type:varchar(200); null"`
	 BmpkExceed             string  `gorm:" column:bmpkExceed;type:varchar(200); null"`
	 MonthlyDebtFreq        int64   `gorm:" column:monthlyDebtFreq;type:int; null"`
	 MonthlyCredFreq        int64   `gorm:" column:monthlyCredFreq;type:int; null"`
	 MonthlyDebtAmt         float32 `gorm:" column:monthlyDebtAmt;type:float; null"`
	 MonthlyCredAmt         float32 `gorm:" column:monthlyCredAmt;type:float; null"`
	 Lbucustypebs2          string  `gorm:" column:lbucustypebs2;type:varchar(200); null"`
	 BirthPlace             string  `gorm:" column:birthPlace;type:varchar(200); null"`
	 Gender                 string  `gorm:" column:gender;type:varchar(200); null"`
	 TitleBeforeName        string  `gorm:" column:titleBeforeName;type:varchar(200); null"`
	 TitleAfterName         string  `gorm:" column:titleAfterName;type:varchar(200); null"`
	 Religion               string  `gorm:" column:religion;type:varchar(200); null"`
	 Resident               string  `gorm:" column:resident;type:varchar(200); null"`
	 MaritalStatus          string  `gorm:" column:maritalStatus;type:varchar(200); null"`
	 SpouseName             string  `gorm:" column:spouseName;type:varchar(200); null"`
	 EducationLevel         string  `gorm:" column:educationLevel;type:varchar(200); null"`
	 Language               string  `gorm:" column:language;type:varchar(200); null"`
	 StaffOfficial          string  `gorm:" column:staffOfficial;type:varchar(200); null"`
	 SidRelation            string  `gorm:" column:sidRelation;type:varchar(200); null"`
	 CustStatus             string  `gorm:" column:custStatus;type:varchar(200); null"`
	 CifBloked              string  `gorm:" column:cifBloked;type:varchar(200); null"`
	 MailingList            string  `gorm:" column:mailingList;type:varchar(200); null"`
	 MailingListStat        string  `gorm:" column:mailingListStat;type:varchar(200); null"`
	 OfficeName             string  `gorm:" column:officeName;type:varchar(200); null"`
	 OfficeAddress          string  `gorm:" column:officeAddress;type:varchar(200); null"`
	 OfficePostcode         string  `gorm:" column:officePostcode;type:varchar(200); null"`
	 SupervisorId           string  `gorm:" column:supervisorId;type:varchar(200); null"`
	 Lastchange             int64   `gorm:" column:lastchange;type:timestamp without time zone; null"`
	 Terminalid             string  `gorm:" column:terminalid;type:varchar(200); null"`
	 IDExpiryDate           int64   `gorm:" column:IDExpiryDate;type:timestamp without time zone; null"`
	 CreationDate           int64   `gorm:" column:CreationDate;type:timestamp without time zone; null"`
	 TanggalLahir           int64   `gorm:" column:TanggalLahir;type:timestamp without time zone; null"`
	 Npwp                   string  `gorm:" column:npwp;type:varchar(200); null"`
	 DomisiliDate           int64   `gorm:" column:domisiliDate;type:timestamp without time zone; null"`
	 Streetname2            string  `gorm:" column:streetname2;type:varchar(200); null"`
	 Pillar2                string  `gorm:" column:pillar2;type:varchar(200); null"`
	 Village2               string  `gorm:" column:village2;type:varchar(200); null"`
	 Subdistrict2           string  `gorm:" column:subdistrict2;type:varchar(200); null"`
	 Citytown2              string  `gorm:" column:citytown2;type:varchar(200); null"`
	 Province2              string  `gorm:" column:province2;type:varchar(200); null"`
	 Postcode2              string  `gorm:" column:Postcode2;type:varchar(200); null"`
	 Streetname3            string  `gorm:" column:streetname3;type:varchar(200); null"`
	 Pillar3                string  `gorm:" column:pillar3;type:varchar(200); null"`
	 Village3               string  `gorm:" column:village3;type:varchar(200); null"`
	 Subdistrict3           string  `gorm:" column:subdistrict3;type:varchar(200); null"`
	 Citytown3              string  `gorm:" column:citytown3;type:varchar(200); null"`
	 Province3              string  `gorm:" column:province3;type:varchar(200); null"`
	 Postcode3              string  `gorm:" column:Postcode3;type:varchar(200); null"`
	 OtherBankAccount       string  `gorm:" column:OtherBankAccount;type:varchar(200); null"`
	 OtherBankName          string  `gorm:" column:OtherBankName;type:varchar(200); null"`
	 OtherBankAccount2      string  `gorm:" column:OtherBankAccount2;type:varchar(200); null"`
	 OtherBankName2         string  `gorm:" column:OtherBankName2;type:varchar(200); null"`
	 HomeOwnership          int64   `gorm:" column:homeOwnership;type:int; null"`
	 Inherited              string  `gorm:" column:inherited;type:varchar(200); null"`
	 EmployeeCode           int64   `gorm:" column:employeeCode;type:int; null"`
	 Staffdesignation       string  `gorm:" column:staffdesignation;type:varchar(200); null"`
	 EmployeeID             string  `gorm:" column:employeeID;type:varchar(200); null"`
	 CustLiability          string  `gorm:" column:custLiability;type:varchar(200); null"`
	 CifPurpose             string  `gorm:" column:cifPurpose;type:varchar(200); null"`
	 Ref                    string  `gorm:" column:ref;type:varchar(200); null"`
	 PhoneOffice            string  `gorm:" column:phoneOffice;type:varchar(200); null"`
	 FaxOffice              string  `gorm:" column:faxOffice;type:varchar(200); null"`
	 EmployeeDate           int64   `gorm:" column:employeeDate;type:timestamp without time zone; null"`
	 OtherBusiness          string  `gorm:" column:otherBusiness;type:varchar(200); null"`
	 BeneficialJob          string  `gorm:" column:beneficialJob;type:varchar(200); null"`
	 RelationCode1          string  `gorm:" column:relationCode1;type:varchar(200); null"`
	 RelationCode2          string  `gorm:" column:relationCode2;type:varchar(200); null"`
	 CustRelated1           string  `gorm:" column:custRelated1;type:varchar(200); null"`
	 CustRelated2           string  `gorm:" column:custRelated2;type:varchar(200); null"`
	 BankRelationType       string  `gorm:" column:bankRelationType;type:varchar(200); null"`
	 NamaCabang             string  `gorm:" column:NamaCabang;type:varchar(100); null"`
	 Msgid                  string  `gorm:" column:msgid;type:varchar(200); null"`
	 Trxid                  string  `gorm:" column:trxid;type:varchar(200); null"`
	 CompanyEntity          string  `gorm:" column:companyEntity;type:varchar(200); null"`
	 EstablishDate          int64   `gorm:" column:establishDate;type:timestamp without time zone; null"`
	 NamePic                string  `gorm:" column:namePic;type:varchar(200); null"`
	 RegPlaceCD             string  `gorm:" column:regPlaceCD;type:varchar(200); null"`
	 Info3                  string  `gorm:" column:info3;type:varchar(200); null"`
	 Info4                  string  `gorm:" column:info4;type:varchar(200); null"`
	 Info5                  string  `gorm:" column:info5;type:varchar(200); null"`
	 Info6                  string  `gorm:" column:info6;type:varchar(200); null"`
	 Info7                  string  `gorm:" column:info7;type:varchar(200); null"`
	 Info8                  string  `gorm:" column:info8;type:varchar(200); null"`
	 Info9                  string  `gorm:" column:info9;type:varchar(200); null"`
	 Info10                 string  `gorm:" column:info10;type:varchar(200); null"`
	 DinNoKYC               string  `gorm:" column:dinNoKYC;type:varchar(200); null"`
	 CompanyNumberCorp      string  `gorm:" column:companyNumberCorp;type:varchar(200); null"`
	 MinistervalidCorp      string  `gorm:" column:ministervalidCorp;type:varchar(200); null"`
	 GopublicCorp           string  `gorm:" column:gopublicCorp;type:varchar(200); null"`
	 DateoccupiedCorp       int64   `gorm:" column:dateoccupiedCorp;type:timestamp without time zone; null"`
	 SidcodeCorp            string  `gorm:" column:sidcodeCorp;type:varchar(200); null"`
	 LhbucodeCorp           string  `gorm:" column:lhbucodeCorp;type:varchar(200); null"`
	 Bamepic1               string  `gorm:" column:namepic1;type:varchar(200); null"`
	 Positionpic1           string  `gorm:" column:positionpic1;type:varchar(200); null"`
	 Phonepic1              string  `gorm:" column:phonepic1;type:varchar(200); null"`
	 Mobilephpic1           string  `gorm:" column:mobilephpic1;type:varchar(200); null"`
	 Emailpic1              string  `gorm:" column:emailpic1;type:varchar(200); null"`
	 Namepic2               string  `gorm:" column:namepic2;type:varchar(200); null"`
	 Positionpic2           string  `gorm:" column:positionpic2;type:varchar(200); null"`
	 Phonepic2              string  `gorm:" column:phonepic2;type:varchar(200); null"`
	 Mobilephpic2           string  `gorm:" column:mobilephpic2;type:varchar(200); null"`
	 Emailpic2              string  `gorm:" column:emailpic2;type:varchar(200); null"`
	 BussNaturecorpinf      string  `gorm:" column:bussNaturecorpinf;type:varchar(200); null"`
	 PrimeBankcorpinf       string  `gorm:" column:primeBankcorpinf;type:varchar(200); null"`
	 BankCodecorpinf        string  `gorm:" column:bankCodecorpinf;type:varchar(200); null"`
	 OwnerStatuscorpinf     string  `gorm:" column:ownerStatuscorpinf;type:varchar(200); null"`
	 Lbucustypebs2CIFBasel2 string  `gorm:" column:lbucustypebs2CIFBasel2;type:varchar(200); null"`
	 AccostsCorp            string  `gorm:" column:accostsCorp;type:varchar(200); null"`
	 TypeCorp               string  `gorm:" column:typeCorp;type:varchar(200); null"`
	 PayrollBatchID         string  `gorm:" column:payrollBatchID;type:varchar(200); null"`
	 SIDBursa               string  `gorm:" column:SIDBursa;type:varchar(200); null"`
	 Name1                  string  `gorm:" column:name1;type:varchar(200); null"`
	 Position1              string  `gorm:" column:position1;type:varchar(200); null"`
	 Address1               string  `gorm:" column:address1;type:varchar(200); null"`
	 Idtype1                string  `gorm:" column:idtype1;type:varchar(200); null"`
	 Idnumber1              string  `gorm:" column:idnumber1;type:varchar(200); null"`
	 Idvalidto1             int64   `gorm:" column:idvalidto1;type:timestamp without time zone; null"`
	 Mobileph1              string  `gorm:" column:mobileph1;type:varchar(200); null"`
	 Gender1                string  `gorm:" column:gender1;type:varchar(200); null"`
	 Npwp1                  string  `gorm:" column:npwp1;type:varchar(200); null"`
	 Citytown1              string  `gorm:" column:citytown1;type:varchar(200); null"`
	 Village1               string  `gorm:" column:village1;type:varchar(200); null"`
	 Subdistrict1           string  `gorm:" column:subdistrict1;type:varchar(200); null"`
	 Name2                  string  `gorm:" column:name2;type:varchar(200); null"`
	 Position2              string  `gorm:" column:position2;type:varchar(200); null"`
	 Address2               string  `gorm:" column:address2;type:varchar(200); null"`
	 Idtype2                string  `gorm:" column:idtype2;type:varchar(200); null"`
	 Idnumber2              string  `gorm:" column:idnumber2;type:varchar(200); null"`
	 Idvalidto2             int64   `gorm:" column:idvalidto2;type:timestamp without time zone; null"`
	 Mobileph2              string  `gorm:" column:mobileph2;type:varchar(200); null"`
	 Gender2                string  `gorm:" column:gender2;type:varchar(200); null"`
	 Npwp2                  string  `gorm:" column:npwp2;type:varchar(200); null"`
	 Name3                  string  `gorm:" column:name3;type:varchar(200); null"`
	 Position3              string  `gorm:" column:position3;type:varchar(200); null"`
	 Address3               string  `gorm:" column:address3;type:varchar(200); null"`
	 Idtype3                string  `gorm:" column:idtype3;type:varchar(200); null"`
	 Idnumber3              string  `gorm:" column:idnumber3;type:varchar(200); null"`
	 Idvalidto3             int64   `gorm:" column:idvalidto3;type:timestamp without time zone; null"`
	 Mobileph3              string  `gorm:" column:mobileph3;type:varchar(200); null"`
	 Gender3                string  `gorm:" column:gender3;type:varchar(200); null"`
	 Npwp3                  string  `gorm:" column:npwp3;type:varchar(200); null"`
	 Name4                  string  `gorm:" column:name4;type:varchar(200); null"`
	 Position4              string  `gorm:" column:position4;type:varchar(200); null"`
	 Address4               string  `gorm:" column:address4;type:varchar(200); null"`
	 Idtype4                string  `gorm:" column:idtype4;type:varchar(200); null"`
	 Idnumber4              string  `gorm:" column:idnumber4;type:varchar(200); null"`
	 Idvalidto4             int64   `gorm:" column:idvalidto4;type:timestamp without time zone; null"`
	 Mobileph4              string  `gorm:" column:mobileph4;type:varchar(200); null"`
	 Gender4                string  `gorm:" column:gender4;type:varchar(200); null"`
	 Npwp4                  string  `gorm:" column:npwp4;type:varchar(200); null"`
	 Citytown4              string  `gorm:" column:citytown4;type:varchar(200); null"`
	 Village4               string  `gorm:" column:village4;type:varchar(200); null"`
	 Subdistrict4           string  `gorm:" column:subdistrict4;type:varchar(200); null"`
	 Citytownp3             string  `gorm:" column:citytownp3;type:varchar(200); null"`
	 Villagep3              string  `gorm:" column:villagep3;type:varchar(200); null"`
	 Subdistrictp3          string  `gorm:" column:subdistrictp3;type:varchar(200); null"`
	 Citytownp2             string  `gorm:" column:citytownp2;type:varchar(200); null"`
	 Villagep2              string  `gorm:" column:villagep2;type:varchar(200); null"`
	 Subdistrictp2          string  `gorm:" column:subdistrictp2;type:varchar(200); null"`
	 Nameo1                 string  `gorm:" column:nameo1;type:varchar(200); null"`
	 Positiono1             string  `gorm:" column:positiono1;type:varchar(200); null"`
	 Addresso1              string  `gorm:" column:addresso1;type:varchar(200); null"`
	 Idtypeo1               string  `gorm:" column:idtypeo1;type:varchar(200); null"`
	 Idnumbero1             string  `gorm:" column:idnumbero1;type:varchar(200); null"`
	 Idvalidtoo1            int64   `gorm:" column:idvalidtoo1;type:timestamp without time zone; null"`
	 Mobilepho1             string  `gorm:" column:mobilepho1;type:varchar(200); null"`
	 SharePercentageo1      string  `gorm:" column:sharePercentageo1;type:decimal(18, 0); null"` //Decimal
	 Gendero1               string  `gorm:" column:gendero1;type:varchar(200); null"`
	 Npwpo1                 string  `gorm:" column:npwpo1;type:varchar(200); null"`
	 Citytowno1             string  `gorm:" column:citytowno1;type:varchar(200); null"`
	 Villageo1              string  `gorm:" column:villageo1;type:varchar(200); null"`
	 Subdistricto1          string  `gorm:" column:subdistricto1;type:varchar(200); null"`
	 Nameo2                 string  `gorm:" column:nameo2;type:varchar(200); null"`
	 Positiono2             string  `gorm:" column:positiono2;type:varchar(200); null"`
	 Addresso2              string  `gorm:" column:addresso2;type:varchar(200); null"`
	 Idtypeo2               string  `gorm:" column:idtypeo2;type:varchar(200); null"`
	 Idnumbero2             string  `gorm:" column:idnumbero2;type:varchar(200); null"`
	 Idvalidtoo2            int64   `gorm:" column:idvalidtoo2;type:timestamp without time zone; null"`
	 Mobilepho2             string  `gorm:" column:mobilepho2;type:varchar(200); null"`
	 SharePercentageo2      string  `gorm:" column:sharePercentageo2;type:decimal(18, 0); null"` //Decimal
	 Gendero2               string  `gorm:" column:gendero2;type:varchar(200); null"`
	 Npwpo2                 string  `gorm:" column:npwpo2;type:varchar(200); null"`
	 Citytowno2             string  `gorm:" column:citytowno2;type:varchar(200); null"`
	 Villageo2              string  `gorm:" column:villageo2;type:varchar(200); null"`
	 Subdistricto2          string  `gorm:" column:subdistricto2;type:varchar(200); null"`
	 Nameo3                 string  `gorm:" column:nameo3;type:varchar(200); null"`
	 Positiono3             string  `gorm:" column:positiono3;type:varchar(200); null"`
	 Addresso3              string  `gorm:" column:addresso3;type:varchar(200); null"`
	 Idtypeo3               string  `gorm:" column:idtypeo3;type:varchar(200); null"`
	 Idnumbero3             string  `gorm:" column:idnumbero3;type:varchar(200); null"`
	 Idvalidtoo3            int64   `gorm:" column:idvalidtoo3;type:timestamp without time zone; null"`
	 Mobilepho3             string  `gorm:" column:mobilepho3;type:varchar(200); null"`
	 SharePercentageo3      string  `gorm:" column:sharePercentageo3;type:decimal(18, 0); null"` //Decimal
	 Gendero3               string  `gorm:" column:gendero3;type:varchar(200); null"`
	 Npwpo3                 string  `gorm:" column:npwpo3;type:varchar(200); null"`
	 Citytowno3             string  `gorm:" column:citytowno3;type:varchar(200); null"`
	 Villageo3              string  `gorm:" column:villageo3;type:varchar(200); null"`
	 Subdistricto3          string  `gorm:" column:subdistricto3;type:varchar(200); null"`
	 Nameo4                 string  `gorm:" column:nameo4;type:varchar(200); null"`
	 Positiono4             string  `gorm:" column:positiono4;type:varchar(200); null"`
	 Addresso4              string  `gorm:" column:addresso4;type:varchar(200); null"`
	 Idtypeo4               string  `gorm:" column:idtypeo4;type:varchar(200); null"`
	 Idnumbero4             string  `gorm:" column:idnumbero4;type:varchar(200); null"`
	 Idvalidtoo4            int64   `gorm:" column:idvalidtoo4;type:timestamp without time zone; null"`
	 Mobilepho4             string  `gorm:" column:mobilepho4;type:varchar(200); null"`
	 SharePercentageo4      string  `gorm:" column:sharePercentageo4;type:decimal(18, 0); null"` //Decimal
	 Gendero4               string  `gorm:" column:gendero4;type:varchar(200); null"`
	 Npwpo4                 string  `gorm:" column:npwpo4;type:varchar(200); null"`
	 Citytowno4             string  `gorm:" column:citytowno4;type:varchar(200); null"`
	 Villageo4              string  `gorm:" column:villageo4;type:varchar(200); null"`
	 Subdistricto4          string  `gorm:" column:subdistricto4;type:varchar(200); null"`
	 Streetname3Corp        string  `gorm:" column:streetname3Corp;type:varchar(200); null"`
	 Citytown3Corp          string  `gorm:" column:citytown3Corp;type:varchar(200); null"`
	 Postcode3Corp          string  `gorm:" column:postcode3Corp;type:varchar(200); null"`
	 TypeBoard              string  `gorm:" column:typeBoard;type:varchar(200); null"`
	 Customer_risk          string  `gorm:" column:customer_risk;type:varchar(200); null"`
	 Eformdate              int64   `gorm:" column:eformdate;type:timestamp without time zone; null"`
	 AppCode                string  `gorm:" column:AppCode;type:varchar(200); null"`
	 VirtualAcc             string  `gorm:" column:VirtualAcc;type:varchar(200); null"`
	 VirtualAccName         string  `gorm:" column:VirtualAccName;type:varchar(200); null"`
	 TransferDate           int64   `gorm:" column:TransferDate;type:timestamp without time zone; null"`
	 HubHukum               string  `gorm:" column:HubHukum;type:varchar(200); null"`
	 HukumLain              string  `gorm:" column:HukumLain;type:varchar(200); null"`
	 OfficePillar           string  `gorm:" column:OfficePillar;type:varchar(200); null"`
	 OfficeVillage          string  `gorm:" column:OfficeVillage;type:varchar(200); null"`
	 OfficeSubdistrict      string  `gorm:" column:OfficeSubdistrict;type:varchar(200); null"`
	 OfficeCityTown         string  `gorm:" column:OfficeCityTown;type:varchar(200); null"`
	 OfficeProvince         string  `gorm:" column:OfficeProvince;type:varchar(200); null"`
	 Jabatan                string  `gorm:" column:Jabatan;type:varchar(200); null"`
	 StatusPegawai          string  `gorm:" column:StatusPegawai;type:varchar(200); null"`
	 JumlahPegawai          string  `gorm:" column:JumlahPegawai;type:varchar(200); null"`
	 NamaDrt                string  `gorm:" column:NamaDrt;type:varchar(200); null"`
	 AlamatDrt              string  `gorm:" column:AlamatDrt;type:varchar(200); null"`
	 PillarDrt              string  `gorm:" column:PillarDrt;type:varchar(200); null"`
	 VillageDrt             string  `gorm:" column:VillageDrt;type:varchar(200); null"`
	 SubdistrictDrt         string  `gorm:" column:SubdistrictDrt;type:varchar(200); null"`
	 CityTownDrt            string  `gorm:" column:CityTownDrt;type:varchar(200); null"`
	 PostCodeDrt            string  `gorm:" column:PostCodeDrt;type:varchar(200); null"`
	 ProvinceDrt            string  `gorm:" column:ProvinceDrt;type:varchar(200); null"`
	 PhoneDrt               string  `gorm:" column:PhoneDrt;type:varchar(200); null"`
	 PonselDrt              string  `gorm:" column:PonselDrt;type:varchar(200); null"`
	 HubunganDrt            string  `gorm:" column:HubunganDrt;type:varchar(200); null"`
	 Flag_source            int64   `gorm:" column:flag_source;type:int; null"`
	 SendEmail              string  `gorm:" column:sendEmail;type:varchar(5); null"`
	 Sales_Code             string  `gorm:" column:Sales_Code;type:varchar(20); null"`
	 Sales_Name             string  `gorm:" column:Sales_Name;type:varchar(100); null"`
	 Sales_Branch_Code      string  `gorm:" column:Sales_Branch_Code;type:varchar(50); null"`
	 Sales_Branch_Name      string  `gorm:" column:Sales_Branch_Name;type:varchar(100); null"`
	 Flag_Create_Sales      int     `gorm:" column:Flag_Create_Sales;type:smallint; null"`
	 Note_tanpa_npwp        string  `gorm:" column:note_tanpa_npwp;type:varchar(100); null"`
	 ReffNo                 string  `gorm:" column:reffNo;type:varchar(10); null"`
	 Data_from              string  `gorm:" column:data_from;type:varchar(50); null"`
	 SIUPNo                 string  `gorm:" column:sIUPNo;type:varchar(200); null"`
	 Commitment             string  `gorm:" column:Commitment;type:varchar(200); null"`
	 ASO_Code               string  `gorm:" column:ASO_Code;type:varchar(200); null"`
	 ASO_Name               string  `gorm:" column:ASO_Name;type:varchar(200); null"`
	 Birthdatep1            string  `gorm:" column:birthdatep1;type:varchar(25); null"`
	 Birthdatep2            string  `gorm:" column:birthdatep2;type:varchar(25); null"`
	 Birthdatep3            string  `gorm:" column:birthdatep3;type:varchar(25); null"`
	 Birthdatep4            string  `gorm:" column:birthdatep4;type:varchar(25); null"`
	 Birthdateo1            string  `gorm:" column:birthdateo1;type:varchar(25); null"`
	 Birthdateo2            string  `gorm:" column:birthdateo2;type:varchar(25); null"`
	 Birthdateo3            string  `gorm:" column:birthdateo3;type:varchar(25); null"`
	 Birthdateo4            string  `gorm:" column:birthdateo4;type:varchar(25); null"`
	 KodeBidangUsaha        string  `gorm:" column:kodeBidangUsaha;type:varchar(200); null"`
	 ReferralCode           string  `gorm:" column:referralCode;type:varchar(100); null"`
	 UsernameOrami          string  `gorm:" column:usernameOrami;type:varchar(100); null"`
	 Birthplacep1           string  `gorm:" column:birthplacep1;type:varchar(200); null"`
	 Birthplacep2           string  `gorm:" column:birthplacep2;type:varchar(200); null"`
	 Birthplacep3           string  `gorm:" column:birthplacep3;type:varchar(200); null"`
	 Birthplacep4           string  `gorm:" column:birthplacep4;type:varchar(200); null"`
	 Birthplaceo1           string  `gorm:" column:birthplaceo1;type:varchar(200); null"`
	 Birthplaceo2           string  `gorm:" column:birthplaceo2;type:varchar(200); null"`
	 Birthplaceo3           string  `gorm:" column:birthplaceo3;type:varchar(200); null"`
	 Birthplaceo4           string  `gorm:" column:birthplaceo4;type:varchar(200); null"`
	 Nationality_code1      string  `gorm:" column:nationality_code1;type:varchar(200); null"`
	 Nationality_code2      string  `gorm:" column:nationality_code2;type:varchar(200); null"`
	 Nationality_code3      string  `gorm:" column:nationality_code3;type:varchar(200); null"`
	 Nationality_code4      string  `gorm:" column:nationality_code4;type:varchar(200); null"`
	 Nationality_codeo1     string  `gorm:" column:nationality_codeo1;type:varchar(200); null"`
	 Nationality_codeo2     string  `gorm:" column:nationality_codeo2;type:varchar(200); null"`
	 Nationality_codeo3     string  `gorm:" column:nationality_codeo3;type:varchar(200); null"`
	 Nationality_codeo4     string  `gorm:" column:nationality_codeo4;type:varchar(200); null"`
	 Country_code1          string  `gorm:" column:country_code1;type:varchar(200); null"`
	 Country_code2          string  `gorm:" column:country_code2;type:varchar(200); null"`
	 Country_code3          string  `gorm:" column:country_code3;type:varchar(200); null"`
	 Country_code4          string  `gorm:" column:country_code4;type:varchar(200); null"`
	 Country_codeo1         string  `gorm:" column:country_codeo1;type:varchar(200); null"`
	 Country_codeo2         string  `gorm:" column:country_codeo2;type:varchar(200); null"`
	 Country_codeo3         string  `gorm:" column:country_codeo3;type:varchar(200); null"`
	 Country_codeo4         string  `gorm:" column:country_codeo4;type:varchar(200); null"`
	 Country_codeAddress1   string  `gorm:" column:country_codeAddress1;type:varchar(200); null"`
	 Country_codeAddress2   string  `gorm:" column:country_codeAddress2;type:varchar(200); null"`
	 Country_codeAddress3   string  `gorm:" column:country_codeAddress3;type:varchar(200); null"`
	 Country_codeAddress4   string  `gorm:" column:country_codeAddress4;type:varchar(200); null"`
	 Country_codeAddresso1  string  `gorm:" column:country_codeAddresso1;type:varchar(200); null"`
	 Country_codeAddresso2  string  `gorm:" column:country_codeAddresso2;type:varchar(200); null"`
	 Country_codeAddresso3  string  `gorm:" column:country_codeAddresso3;type:varchar(200); null"`
	 Country_codeAddresso4  string  `gorm:" column:country_codeAddresso4;type:varchar(200); null"`
	 Country_codeIdentity1  string  `gorm:" column:country_codeIdentity1;type:varchar(200); null"`
	 Country_codeIdentity2  string  `gorm:" column:country_codeIdentity2;type:varchar(200); null"`
	 Country_codeIdentity3  string  `gorm:" column:country_codeIdentity3;type:varchar(200); null"`
	 Country_codeIdentity4  string  `gorm:" column:country_codeIdentity4;type:varchar(200); null"`
	 Country_codeIdentityo1 string  `gorm:" column:country_codeIdentityo1;type:varchar(200); null"`
	 Country_codeIdentityo2 string  `gorm:" column:country_codeIdentityo2;type:varchar(200); null"`
	 Country_codeIdentityo3 string  `gorm:" column:country_codeIdentityo3;type:varchar(200); null"`
	 Country_codeIdentityo4 string  `gorm:" column:country_codeIdentityo4;type:varchar(200); null"`
	 Source1                string  `gorm:" column:source1;type:varchar(200); null"`
	 Source2                string  `gorm:" column:source2;type:varchar(200); null"`
	 Source3                string  `gorm:" column:source3;type:varchar(200); null"`
	 Source4                string  `gorm:" column:source4;type:varchar(200); null"`
	 Sourceo1               string  `gorm:" column:sourceo1;type:varchar(200); null"`
	 Sourceo2               string  `gorm:" column:sourceo2;type:varchar(200); null"`
	 Sourceo3               string  `gorm:" column:sourceo3;type:varchar(200); null"`
	 Sourceo4               string  `gorm:" column:sourceo4;type:varchar(200); null"`
	 Pekerjaan1             string  `gorm:" column:pekerjaan1;type:varchar(200); null"`
	 Pekerjaan2             string  `gorm:" column:pekerjaan2;type:varchar(200); null"`
	 Pekerjaan3             string  `gorm:" column:pekerjaan3;type:varchar(200); null"`
	 Pekerjaan4             string  `gorm:" column:pekerjaan4;type:varchar(200); null"`
	 Pekerjaano1            string  `gorm:" column:pekerjaano1;type:varchar(200); null"`
	 Pekerjaano2            string  `gorm:" column:pekerjaano2;type:varchar(200); null"`
	 Pekerjaano3            string  `gorm:" column:pekerjaano3;type:varchar(200); null"`
	 Pekerjaano4            string  `gorm:" column:pekerjaano4;type:varchar(200); null"`
	 Provincen1             string  `gorm:" column:provincen1;type:varchar(200); null"`
	 Provincen2             string  `gorm:" column:provincen2;type:varchar(200); null"`
	 Provincen3             string  `gorm:" column:provincen3;type:varchar(200); null"`
	 Provincen4             string  `gorm:" column:provincen4;type:varchar(200); null"`
	 Provinceo1             string  `gorm:" column:provinceo1;type:varchar(200); null"`
	 Provinceo2             string  `gorm:" column:provinceo2;type:varchar(200); null"`
	 Provinceo3             string  `gorm:" column:provinceo3;type:varchar(200); null"`
	 Provinceo4             string  `gorm:" column:provinceo4;type:varchar(200); null"`
	 Post_code1             string  `gorm:" column:post_code1;type:varchar(200); null"`
	 Post_code2             string  `gorm:" column:post_code2;type:varchar(200); null"`
	 Post_code3             string  `gorm:" column:post_code3;type:varchar(200); null"`
	 Post_code4             string  `gorm:" column:post_code4;type:varchar(200); null"`
	 Post_codeo1            string  `gorm:" column:post_codeo1;type:varchar(200); null"`
	 Post_codeo2            string  `gorm:" column:post_codeo2;type:varchar(200); null"`
	 Post_codeo3            string  `gorm:" column:post_codeo3;type:varchar(200); null"`
	 Post_codeo4            string  `gorm:" column:post_codeo4;type:varchar(200); null"`
	 FaxNo                  string  `gorm:" column:faxNo;type:varchar(50); null"`
	 IncomeLevel            int64   `gorm:" column:IncomeLevel;type:int; null"`
	 ContactPhoneOffice     string  `gorm:" column:contactPhoneOffice;type:varchar(50); null"`
	 CustRisk               string  `gorm:" column:CustRisk;type:varchar(5); null"`
	 Street2                string  `gorm:" column:Street2;type:varchar(50); null"`
	 PrvPhyAdd2             string  `gorm:" column:PrvPhyAdd2;type:varchar(50); null"`
	 PoBoxNo2               string  `gorm:" column:PoBoxNo2;type:varchar(50); null"`
	 CRN1                   string  `gorm:" column:CRN1;type:varchar(20); null"`
	 CRN2                   string  `gorm:" column:CRN2;type:varchar(20); null"`
	 CRN3                   string  `gorm:" column:CRN3;type:varchar(50); null"`
	 Info3CD                string  `gorm:" column:info3CD;type:varchar(200); null"`
	 IsPEP                  string  `gorm:" column:IsPEP;type:varchar(10); null"`
	 CrossSell              string  `gorm:" column:crossSell;type:varchar(10); null"`
	 Full_name              string  `gorm:" column:full_name;type:varchar(200); null"`
	 PisahHarta             string  `gorm:" column:pisahHarta;type:varchar(10); null"`
	 BirthDatePasangan      string  `gorm:" column:birthDatePasangan;type:varchar(25); null"`
	 NikPasangan            string  `gorm:" column:nikPasangan;type:varchar(200); null"`
	 JumlahTanggungan       string  `gorm:" column:jumlahTanggungan;type:varchar(10); null"`
	 IsWorking              string  `gorm:" column:isWorking;type:varchar(10); null"`
	 GroupDebitur           string  `gorm:" column:groupDebitur;type:varchar(200); null"`
	 RatingDebitur          string  `gorm:" column:ratingDebitur;type:varchar(25); null"`
	 LembagaPemeringkat     string  `gorm:" column:lembagaPemeringkat;type:varchar(25); null"`
	 TanggalPemeringkat     string  `gorm:" column:tanggalPemeringkat;type:varchar(25); null"`
	 JenisKartuPasangan     string  `gorm:" column:jenisKartuPasangan;type:varchar(200); null"`
	 Kseisid                string  `gorm:" column:kseisid;type:varchar(200); null"`
	 FirstEstablishCode     string  `gorm:" column:FirstEstablishCode;type:varchar(200); null"`
	 FlagBenefOwn           string  `gorm:" column:FlagBenefOwn;type:varchar(10); null"`
	 NasabahPrioritas       string  `gorm:" column:nasabahPrioritas;type:varchar(20); null"`
	 FirstEstablishDate     int64   `gorm:" column:FirstEstablishDate;type:timestamp without time zone; null"`
 }